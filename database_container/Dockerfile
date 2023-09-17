FROM oryd/hydra:v1.10.6

# The system secret can only be set against a fresh database. Key rotation is not supported. This
# secret is used to encrypt the database and needs to be set to the same value every time the process (re-)starts.
# You can use /dev/urandom to generate a secret. But make sure that the secret must be the same anytime you define it.
# You could, for example, store the value somewhere.
ENV SECRETS_SYSTEM=$("export LC_CTYPE=C; cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1")

# Alternatively you can obviously just set a secret:
# export SECRETS_SYSTEM=this_needs_to_be_the_same_always_and_also_very_$3cuR3-._

# The database url points us at the postgres instance. This could also be an ephemeral in-memory database (`export DSN=memory`)
# or a MySQL URI.
ENV DSN=postgres://hydra:secret@host.docker.internal:5432/hydra?sslmode=disable
