package gray

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/oldthreefeng/mycli/k8s"
	"github.com/oldthreefeng/mycli/utils"
	"github.com/wonderivan/logger"
	"k8s.io/client-go/kubernetes"
)

type project struct {
	Name    string  `json:"name"`
	Version string	`json:"version"`
	Image   string 	`json:"image"`
	Syncd   bool	`json:"syncd"`
	Client *kubernetes.Clientset
	debug bool		
}

const (
	defaultNamespace = "prod"
	defaultName      = "be-learntask-service"
)

func (p *project) sync(p2 project) {
	if p.Image == p2.Image {
		p.Syncd = true
		bytes, err := json.Marshal(p)
		if err != nil {
			logger.Error("marshal json file error")
			return
		}
		fmt.Println(string(bytes))
	} else {
		if p.debug {
			logger.Warn("version: %s , v1 deployment: %s , v1image: %s, you need to update your v2 deployment : %s, , cause v2image is: %s", 
			p.Version, p.Name,p.Image, p2.Name, p2.Image)
			return
		}
		err := k8s.SetDpImage(p.Client, p2.Name, defaultNamespace, p.Image)
		if err != nil {
			logger.Error("set v2deployment image error: ", err)
			return 
		}
		logger.Info("v2 deployment [%s] has already syncd v1 image: %s",p2.Name, p.Image )
		p.Syncd = true
	}
}

func Sync(debug bool) {
	syncImage(debug)
}

func syncImage(debug bool) {
	ps, err := getprojectList(debug)
	if err != nil {
		utils.ProcessError(err)
	}
	k8sClient, err := k8s.NewClient(nil)
	if err != nil {
		utils.ProcessError(err)
		return
	}
	for _ , v := range ps {
		var p1 project
		p1.Client = k8sClient
		p1.Name = v.Name[:len(v.Name)-3]
		v1Image, version, err := k8s.GetImageByDpName(k8sClient, p1.Name, defaultNamespace)
		if err != nil {
			utils.ProcessError(err)
			return
		}
		p1.Image = v1Image
		p1.Version = version
		p1.debug = v.debug
		p1.sync(v)
	}
}

func getprojectList(debug bool) ([]project ,error){
	k8sClinet, err := k8s.NewClient(nil)
	if err != nil {
		return nil, utils.ProcessError(err)
	}

	// get deployments
	dps, err := k8s.GetDps(k8sClinet, defaultNamespace)
	if err != nil {
		return nil, utils.ProcessError(err)
	}

	var ps []project
	// get deployments name, version, image
	for _, dep := range dps.Items {
		// 主要是同步v2版本， 有v2则必有v1.
		if strings.Contains(dep.GetName(), "v2") && strings.Contains(dep.GetName(), "be"){
			// get deployment containers
			var p project
			
			// 忽略灰度为0 的实例。
			// if *dep.Spec.Replicas == 0 {
			// 	continue
			// }

			for _, ctr := range dep.Spec.Template.Spec.Containers {
				// get our main container
				if strings.Contains(ctr.Name, dep.GetName()) {
					// get env
					for _, en := range ctr.Env {
						if en.Name == "VERSION" {
							p.Version = en.Value
							break
						}
					}
					// get Image
					p.Image = ctr.Image
					p.debug = debug
					p.Name = dep.GetName()
					ps = append(ps, p)
					break
				}
			}		
		}
	}
	return ps, nil
}


