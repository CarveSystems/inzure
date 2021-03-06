_cli_zsh_autocomplete() {
    local -a completions
    local -a completions_with_descriptions
    local -a response
    response=("${(@f)$( env COMP_WORDS="${words[*]}" \
        COMP_CWORD=$((CURRENT-1)) \
        _INZURE_COMPLETE="complete_zsh" \
        inzure )}")

    for key descr in ${(kv)response}; do
        if [[ "$descr" == "_" ]]; then
            completions+=("$key")
        else
            completions_with_descriptions+=("$key":"$descr")
        fi
    done

    if [ -n "$completions_with_descriptions" ]; then
        _describe -V unsorted completions_with_descriptions -U -Q
    fi

    if [ -n "$completions" ]; then
        compadd -S "" -U -V unsorted -Q -a completions
    fi
    compstate[insert]="automenu"
}

compdef _cli_zsh_autocomplete inzure
