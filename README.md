## fcm_sender

##### сервис для отправки пуш-уведомлений

#### запукс fcm_sender как service daemon

скопировать `bin/fcm_sender.service` в `/etc/systemd/system`

```bash
    systemctl enable fcm_sender.service
    systemctl start fcm_sender.service
```