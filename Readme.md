#  tfa
tfa is 2fa cli tool that aim to help you to generate 2fa code on CI/CD pipelines.

You can provide secret with stdin or flag.

## Flag usage
```shell
tfa -key="GEZDGNBV"
```

> Your 2fa code: 160014

## Stdin usage
```shell
echo -n "GEZDGNBV" | tfa
```

> Your 2fa code: 160014