export ENV=local
cdklocal bootstrap -v --profile default
cdklocal synth
cdklocal deploy --require-approval never -v --profile default