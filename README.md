# sync2snipe
sync different data sources to snipe-it

## Configuration
`python3 -m pip install -r requirements.txt`

## Authentication
By default, this will use [keyring](https://github.com/jaraco/keyring) to store your sensitive credentials in your keychain.

If you wish to use this program in a CI envrionment, or elsewhere, where the keychain is not a feasible option you can either instantiating the ToSnipe class with `env_vars=True` or pass the command line flag `--use-env-vars`.

## Sync Jamf items to Snipe

## Usage
```
python3 jamf2.py -h
usage: jamf2.py [-h] [-d] [--disable-requests-logging] [--do_not_verify_ssl] [--dryrun] [-f] [-l] [-lf LOG_FILE] [-rt] [-s SETTINGS_FILE] [-v] [--auto_incrementing] [--do_not_update_jamf] [-u | -ui | -uf] [-uns] [-m | -c]

options:
  -h, --help            show this help message and exit
  -d, --debug           Sets logging to include additional DEBUG messages.
  --disable-requests-logging
                        In debug logging disables the requests library logs
  --do_not_verify_ssl   Skips SSL verification for all requests. Helpful when you use self-signed certificate.
  --dryrun              This checks your config and tries to contact both the JAMFPro and Snipe-it instances, but exits before updating or syncing any assets.
  -f, --force           Updates the Snipe asset with information every time, despite what the timestamps indicate.
  -l, --log-to-file     Output log results to file.
  -lf LOG_FILE, --log-file LOG_FILE
                        location of log file. defaults to sync2snipe-$date.log.
  -rt, --run-tests      Run startup tests to see if hosts are reachable.
  -s SETTINGS_FILE, --settings-file SETTINGS_FILE
                        location of settings file. defaults to settings.json.
  -v, --verbose         Sets the logging level to INFO and gives you a better idea of what the script is doing.
  --auto_incrementing   You can use this if you have auto-incrementing enabled in your snipe instance to utilize that instead of adding the Jamf ID for the asset tag.
  --do_not_update_jamf  Does not update Jamf with the asset tags stored in Snipe.
  -u, --users           Checks out the item to the current user in Jamf if it's not already deployed
  -ui, --users_inverse  Checks out the item to the current user in Jamf if it's already deployed
  -uf, --users_force    Checks out the item to the user specified in Jamf no matter what
  -uns, --users_no_search
                        Doesn't search for any users if the specified fields in Jamf and Snipe don't match. (case insensitive)
  -m, --mobiles         Runs against the Jamf mobiles endpoint only.
  -c, --computers       Runs against the Jamf computers endpoint only.
```

If you pass the `--dryrun` flag the program will run through each item logging what it _would_ be doing.

## Credit
* This project would not be possible without the wonderful folks at Snipe-IT
Without jamf2snipe this project probably wouldnt exist
* The jamf python [library](https://github.com/univ-of-utah-marriott-library-apple/python-jamf) was the source for all the good bits of the jamf classes.
