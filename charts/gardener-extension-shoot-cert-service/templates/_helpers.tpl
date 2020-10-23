{{- define "name" -}}
gardener-extension-shoot-cert-service
{{- end -}}

{{- define "certconfig" -}}
---
apiVersion: shoot-cert-service.extensions.config.gardener.cloud/v1alpha1
kind: Configuration
issuerName: {{ required ".Values.certificateConfig.defaultIssuer.name is required" .Values.certificateConfig.defaultIssuer.name }}
restrictIssuer: {{ required ".Values.certificateConfig.defaultIssuer.restricted is required" .Values.certificateConfig.defaultIssuer.restricted }}
acme:
  email: {{ required ".Values.certificateConfig.defaultIssuer.acme.email is required" .Values.certificateConfig.defaultIssuer.acme.email }}
  server: {{ required ".Values.certificateConfig.defaultIssuer.acme.server is required" .Values.certificateConfig.defaultIssuer.acme.server }}
  {{- if .Values.certificateConfig.defaultIssuer.acme.privateKey }}
  privateKey: |
{{ .Values.certificateConfig.defaultIssuer.acme.privateKey | trim | indent 4 }}
  {{- end }}
{{- end }}

{{-  define "image" -}}
  {{- if hasPrefix "sha256:" .Values.image.tag }}
  {{- printf "%s@%s" .Values.image.repository .Values.image.tag }}
  {{- else }}
  {{- printf "%s:%s" .Values.image.repository .Values.image.tag }}
  {{- end }}
{{- end }}

{{- define "priorityclassversion" -}}
{{- if semverCompare ">= 1.14-0" .Capabilities.KubeVersion.GitVersion -}}
scheduling.k8s.io/v1
{{- else if semverCompare ">= 1.11-0" .Capabilities.KubeVersion.GitVersion -}}
scheduling.k8s.io/v1beta1
{{- else -}}
scheduling.k8s.io/v1alpha1
{{- end -}}
{{- end -}}