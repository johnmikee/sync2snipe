# Where should we write the GIF?
Output keyring.gif

# Set up a 1200x600 terminal with 46px font.
Set FontSize 46
Set Width 1800
Set Height 600

# Show the Output
Show

# Type a command in the terminal.
Type "python3 j2s.py --dryrun --verbose"
Enter

Sleep 500ms

# Enter the keyring prompt
Type "ABC123"

# Submit
Enter

Sleep 500ms

# Enter the keyring prompt
Type "ABC1234"

# Submit
Enter

# Hide the errors
Hide

Sleep 15
