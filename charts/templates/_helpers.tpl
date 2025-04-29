{{- define "simplebank.name" -}}
simplebank
{{- end }}

{{- define "simplebank.fullname" -}}
{{ .Release.Name }}-{{ include "simplebank.name" . }}
{{- end }}
