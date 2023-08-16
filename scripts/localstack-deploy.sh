cdklocal synth
cdklocal bootstrap -v
ENV=local cdklocal deploy --require-approval never -v