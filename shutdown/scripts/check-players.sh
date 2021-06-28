#!/bin/bash
PLAYERS=$(factorio players-online)
source /opt/NO_PLAYERS_PRESENT
echo "Players:"
echo $PLAYERS
if [ -z "$PLAYERS" ]
then
  # No Players in the server
  echo "No players present"
  export NO_PLAYERS_PRESENT=$(($NO_PLAYERS_PRESENT+1))
  gcloud beta logging write no_players_check_log "No players for $NO_PLAYERS_PRESENT checks" --severity=NOTICE
else
  echo "Players present"
  export NO_PLAYERS_PRESENT=0
fi
if [ $NO_PLAYERS_PRESENT -ge 3 ]; then
  # No players in 3 checks, shutdown server
  echo "No players for 3 checks... Shutting down"
  gcloud beta logging write no_players_check_log "SHUTTING DOWN" --severity=NOTICE
  factorio stop
  shutdown -h now
fi
printf "export NO_PLAYERS_PRESENT=%s" $NO_PLAYERS_PRESENT > /opt/NO_PLAYERS_PRESENT