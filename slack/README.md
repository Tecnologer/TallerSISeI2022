https://api.slack.com/apps/
https://api.slack.com/methods/chat.postMessage
https://api.slack.com/methods/conversations.history
https://api.slack.com/methods/conversations.list

```json
{
    "display_information": {
        "name": "SiSei"
    },
    "features": {
        "app_home": {
            "home_tab_enabled": true,
            "messages_tab_enabled": true,
            "messages_tab_read_only_enabled": false
        },
        "bot_user": {
            "display_name": "SiSei",
            "always_online": true
        }
    },
    "oauth_config": {
        "scopes": {
            "bot": [
                "chat:write",
                "channels:read",
                "groups:read",
                "im:read",
                "mpim:read"
            ]
        }
    },
    "settings": {
        "org_deploy_enabled": false,
        "socket_mode_enabled": false,
        "token_rotation_enabled": false
    }
}
```