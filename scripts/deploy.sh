rm -rd cdk.out
cdk synth
cdk deploy --require-approval never --outputs-file .deployed-env