package templates

const SelectionFilterPlaceholderText = "type to filter choices"
const SelectionInputPlaceholderText = "cannot be empty"
const SelectionSelectTemplate = `
{{- if .Prompt -}}
  {{ Bold .Prompt }}
{{ end -}}

{{ if .IsFiltered }}
  {{- print "Filter: " .FilterInput }}
{{ end }}

{{- if not (eq (len .Choices) 0)}}
{{- print "Command: " (index .Choices $.SelectedIndex).Value "\n"}}
{{- end }}

{{- range  $i, $choice := .Choices }}
  {{- if or (IsScrollUpHintPosition $i) (IsScrollDownHintPosition $i) }}
	{{- "• " -}}
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
const SelectionResultTemplate = `
{{- print (Final .FinalChoice) "\n" -}}
`
