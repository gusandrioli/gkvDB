#!/bin/sh
crontab -l > new_crontab
echo "*/1 * * * * ~/Documents/PP/gkvDB/bin/cleaner >> ~/Documents/PP/gkvDB/gkvDB.log 2>&1" >> new_crontab
crontab new_crontab
rm new_crontab