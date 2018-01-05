package main

import (
        "gcodebase/notify"
)

func main() {
        notify := notify.GetSlackNotificationFromParam(&notify.NotifyParam{
                SlackToken:"xoxb-49744114386-4DcUp3ftGubP4OMoZv8KKn9W",
                SlackTopic:"XXXXX",
                SlackChannelName:"glorydns"})
        notify.Notify("sourceXX","msg XXXX")

}
