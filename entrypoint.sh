#!/bin/bash

set -euo pipefail

# Initialize
if [[ "${1:-run}" == init ]]; then

    # Initialize pass
    gpg --generate-key --batch /protonmail/gpgparams
    pass init pass-key

    # Login
    protonmail-bridge --cli

else

    # Start fake mail auth server
    fakeauth &

    # nginx will proxy the connection and make it appear to come from 127.0.0.1
    # ProtonMail Bridge currently expects that.
    nginx

    # Start protonmail
    # Fake a terminal, so it does not quit because of EOF...
    rm -f faketty
    mkfifo faketty
    cat faketty | protonmail-bridge --cli

fi
