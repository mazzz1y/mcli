package prompt

const selectionSelectTemplate = `
{{- if .Prompt -}}
  {{ Bold .Prompt }}
{{ end -}}

{{ if .IsFiltered }}
  {{- print FilterPromptStyle ": " .FilterInput }}
{{ end }}

{{- if not (eq (len .Choices) 0)}}
  {{- print CommandPromptStyle ": " (CommandStyle (index .Choices $.SelectedIndex)) "\n"}}
{{- end }}

{{- range  $i, $choice := .Choices }}
  {{- if or (IsScrollUpHintPosition $i) (IsScrollDownHintPosition $i) }}
    {{- print UpDownSymbolStyle " " -}}
  {{- else -}}
    {{- "  " -}}
  {{- end -}}

  {{- if eq $.SelectedIndex $i }}
    {{- print (Foreground "32" (Bold "~ ")) (Selected $choice) "\n" }}
  {{- else }}
	  {{- print "  " (Unselected $choice) "\n" }}
  {{- end }}
{{- end}}
`
const selectionResultTemplate = `
{{- print FinalSymbolStyle " " (Final .FinalChoice) "\n" -}}
`
