export ENV=local
cdklocal synth
cdklocal bootstrap -v
cdklocal deploy --require-approval never -v