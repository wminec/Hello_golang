# OCP Pod State Watcher
This app is OCP Pods Watcher. It logs when the Phase of a Pod in a specific namespace becomes Running, Failed, or Succeeded.
There are two types of logs:
1. Container stdout (always logged).
2. Email (if configured, alerts can be received via email).

The following environment variables can be configured:
1. NAMESPACE: The namespace where the Pods to be monitored are located.
2. LABEL_KEYS: The label information of the Pods to be monitored. It is separated by commas (,) and uses OR operation.
3. TZ: The timezone of the Pods.
4. EMAIL: Whether to receive alerts via email.
5. EMAIL_FROM: The "from" field of the email.
6. EMAIL_TO: The email address to receive alerts.
7. SMTP_SERVER: The SMTP server address (e.g., smtp.google.com).
8. SMTP_PORT: The port of the SMTP server.
9. SMTP_USER: The user information to connect to the SMTP server (recommended to be set as a secret).
10. SMTP_PASSWORD: The password of the user to connect to the SMTP server (recommended to be set as a secret).