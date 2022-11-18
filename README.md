# sync2snipe
sync different data sources to snipe-it

## Configuration
`python3 -m pip install -r requirements.txt`

## Authentication
By default, this will use [keyring](https://github.com/jaraco/keyring) to store your sensitive credentials in your keychain.

If you wish to use this program in a CI envrionment, or elsewhere, where the keychain is not a feasible option you can either instantiating the Jamf2Snipe class with `env_vars=True` or pass the command line flag `--use-env-vars`.

## Sync Jamf items to Snipe

## Usage

## Credit
* This project would not be possible without the wonderful folks at Snipe-IT
Without jamf2snipe this project probably wouldnt exist
* The jamf python [library](https://github.com/univ-of-utah-marriott-library-apple/python-jamf) was the source for all the good bits of the jamf classes.
