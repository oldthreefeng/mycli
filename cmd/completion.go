/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wonderivan/logger"
	"io"
	"os"
)

const defaultCopyRight = `
# Copyright © 2020 mycli.

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
`

var (
	completionShells = map[string]func(out io.Writer, boilerPlate string, cmd *cobra.Command) error{
		"bash": runCompletionBash,
		"zsh": runCompletionZsh,
	}
	completionExample = `
	# Installing bash completion on macOS using homebrew
		## If running Bash 3.2 included with macOS
		    brew install bash-completion
		## or, if running Bash 4.1+
		    brew install bash-completion@2
		## If you've installed via other means, you may need add the completion to your completion directory
		    mycli completion bash > $(brew --prefix)/etc/bash_completion.d/mycli

	# Load the mycli completion code for zsh[1] into the current shell
		    source <(mycli completion zsh)
	# Set the mycli completion code for zsh[1] to autoload on startup
		    mycli completion zsh > "${fpath[1]}/_mycli"

	# Installing bash completion on Linux
		## If bash-completion is not installed on Linux, please install the 'bash-completion' package
		## via your distribution's package manager.
		## Load the mycli completion code for bash into the current shell
		    source <(mycli completion bash)
		## Write bash completion code to a file and source if from .bash_profile
		    mycli completion bash > ~/.mycli/completion.bash.inc
		    printf "
		      # mycli shell completion
			  source '$HOME/.mycli/completion.bash.inc'
		      " >> $HOME/.bash_profile
		      source $HOME/.bash_profile
`
)

func init() {
	rootCmd.AddCommand(NewCmdCompletion(os.Stdout, ""))
}

func NewCmdCompletion(out io.Writer, boilerPlate string) *cobra.Command {
	shells := []string{}
	for s := range completionShells {
		shells = append(shells, s)
	}

	cmd := &cobra.Command{
		Use:                   "completion bash",
		DisableFlagsInUseLine: true,
		Short:                 "Output shell completion code for the specified shell (bash or zsh)",
		Example: completionExample,
		Run: func(cmd *cobra.Command, args []string) {
			err := RunCompletion(out, boilerPlate, cmd, args)
			if err != nil {
				logger.Error(err)
				return
			}
		},
		ValidArgs: shells,
	}
	return cmd
}

func RunCompletion(out io.Writer, copyRight string, cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("shell not specified")
	}
	if len(args) > 1 {
		return fmt.Errorf("too many arguments, expected only the shell type")
	}
	run, found := completionShells[args[0]]
	if !found {
		return fmt.Errorf("unsupported shell type %q", args[0])
	}

	return run(out, copyRight, cmd.Parent())
}

func runCompletionBash(out io.Writer, copyRight string, mycli *cobra.Command) error {
	if len(copyRight) == 0 {
		copyRight = defaultCopyRight
	}
	if _, err := out.Write([]byte(copyRight)); err != nil {
		return err
	}

	return mycli.GenBashCompletion(out)
}


func runCompletionZsh(out io.Writer, copyRight string, mycli *cobra.Command) error {
	zshHead := "#compdef mycli\n"

	out.Write([]byte(zshHead))

	if len(copyRight) == 0 {
		copyRight = defaultCopyRight
	}
	if _, err := out.Write([]byte(copyRight)); err != nil {
		return err
	}

	zshInitialization := `
__mycli_bash_source() {
	alias shopt=':'
	emulate -L sh
	setopt kshglob noshglob braceexpand
	source "$@"
}
__mycli_type() {
	# -t is not supported by zsh
	if [ "$1" == "-t" ]; then
		shift
		# fake Bash 4 to disable "complete -o nospace". Instead
		# "compopt +-o nospace" is used in the code to toggle trailing
		# spaces. We don't support that, but leave trailing spaces on
		# all the time
		if [ "$1" = "__mycli_compopt" ]; then
			echo builtin
			return 0
		fi
	fi
	type "$@"
}
__mycli_compgen() {
	local completions w
	completions=( $(compgen "$@") ) || return $?
	# filter by given word as prefix
	while [[ "$1" = -* && "$1" != -- ]]; do
		shift
		shift
	done
	if [[ "$1" == -- ]]; then
		shift
	fi
	for w in "${completions[@]}"; do
		if [[ "${w}" = "$1"* ]]; then
			echo "${w}"
		fi
	done
}
__mycli_compopt() {
	true # don't do anything. Not supported by bashcompinit in zsh
}
__mycli_ltrim_colon_completions()
{
	if [[ "$1" == *:* && "$COMP_WORDBREAKS" == *:* ]]; then
		# Remove colon-word prefix from COMPREPLY items
		local colon_word=${1%${1##*:}}
		local i=${#COMPREPLY[*]}
		while [[ $((--i)) -ge 0 ]]; do
			COMPREPLY[$i]=${COMPREPLY[$i]#"$colon_word"}
		done
	fi
}
__mycli_get_comp_words_by_ref() {
	cur="${COMP_WORDS[COMP_CWORD]}"
	prev="${COMP_WORDS[${COMP_CWORD}-1]}"
	words=("${COMP_WORDS[@]}")
	cword=("${COMP_CWORD[@]}")
}
__mycli_filedir() {
	# Don't need to do anything here.
	# Otherwise we will get trailing space without "compopt -o nospace"
	true
}
autoload -U +X bashcompinit && bashcompinit
# use word boundary patterns for BSD or GNU sed
LWORD='[[:<:]]'
RWORD='[[:>:]]'
if sed --help 2>&1 | grep -q 'GNU\|BusyBox'; then
	LWORD='\<'
	RWORD='\>'
fi
__mycli_convert_bash_to_zsh() {
	sed \
	-e 's/declare -F/whence -w/' \
	-e 's/_get_comp_words_by_ref "\$@"/_get_comp_words_by_ref "\$*"/' \
	-e 's/local \([a-zA-Z0-9_]*\)=/local \1; \1=/' \
	-e 's/flags+=("\(--.*\)=")/flags+=("\1"); two_word_flags+=("\1")/' \
	-e 's/must_have_one_flag+=("\(--.*\)=")/must_have_one_flag+=("\1")/' \
	-e "s/${LWORD}_filedir${RWORD}/__mycli_filedir/g" \
	-e "s/${LWORD}_get_comp_words_by_ref${RWORD}/__mycli_get_comp_words_by_ref/g" \
	-e "s/${LWORD}__ltrim_colon_completions${RWORD}/__mycli_ltrim_colon_completions/g" \
	-e "s/${LWORD}compgen${RWORD}/__mycli_compgen/g" \
	-e "s/${LWORD}compopt${RWORD}/__mycli_compopt/g" \
	-e "s/${LWORD}declare${RWORD}/builtin declare/g" \
	-e "s/\\\$(type${RWORD}/\$(__mycli_type/g" \
	<<'BASH_COMPLETION_EOF'
`
	out.Write([]byte(zshInitialization))

	buf := new(bytes.Buffer)
	mycli.GenBashCompletion(buf)
	out.Write(buf.Bytes())

	zshTail := `
BASH_COMPLETION_EOF
}
__mycli_bash_source <(__mycli_convert_bash_to_zsh)
`
	out.Write([]byte(zshTail))
	return nil
}