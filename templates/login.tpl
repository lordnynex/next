{{ define "login" }}
{{ template "layout/header" . }}
<p>
  Click the link below to login into your account.
  This link will expire in 15 minutes and can only be used once.
</p>
<a href="{{ .link }}">Log in</a>
{{ template "layout/footer" . }}
{{ end }}
