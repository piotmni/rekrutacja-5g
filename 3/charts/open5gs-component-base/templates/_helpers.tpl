{{/*
Set name variable based on component name
*/}}
{{- define "open5gs-component-base.name" -}}
{{- printf "open5gs-%s" .Values.componentName }}
{{- end }}


{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "open5gs-component-base.chart" -}}
{{- printf "open5gs-%s-%s" .Values.componentName .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}


{{- define "open5gs-component-base.labels" -}}
helm.sh/chart: {{ include "open5gs-component-base.chart" . }}
{{ include "open5gs-component-base.selectorLabels" . }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "open5gs-component-base.selectorLabels" -}}
app.kubernetes.io/name: {{ include "open5gs-component-base.name" . }}
app.kubernetes.io/instance: open5gs-{{ .Values.componentName }}
{{- end }}