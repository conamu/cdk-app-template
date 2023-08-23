export ENV=local
cdklocal bootstrap -v
cdklocal synth
cdklocal deploy --require-approval never -v