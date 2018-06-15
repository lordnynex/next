{{ define "login" }}
{{ template "layout/header" . }}
<p>
  Click the link below to login into your account.
  This link will expire in 10 minutes and can only be used once.
</p>
<a href="{{ .loginLink }}">Log in</a>
<p>
  Or just type the disposable code right in the app: {{ .code }}
</p>
{{ template "layout/footer" . }}
{{ end }}
