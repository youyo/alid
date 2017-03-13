# alid

alid is *A*mazon *L*inux latest ami *ID*.

## Description

## Usage

First, setup AWS Access key and Secret access key.

```
$ alid
ami-b3d2abd4
```

Specific region.

```
$ alid -r ap-northeast-1
ami-b3d2abd4
```

AMI all information.

```
$ alid -all
{
  Architecture: "x86_64",
  BlockDeviceMappings: [{
      DeviceName: "/dev/xvda",
      Ebs: {
        DeleteOnTermination: true,
        Encrypted: false,
        SnapshotId: "snap-0e57c6c5c836f4bf3",
        VolumeSize: 8,
        VolumeType: "standard"
      }
    }],
  CreationDate: "2017-01-20T23:39:50.000Z",
  Description: "Amazon Linux AMI 2016.09.1.20170119 x86_64 HVM EBS",
  EnaSupport: true,
  Hypervisor: "xen",
  ImageId: "ami-b3d2abd4",
  ImageLocation: "amazon/amzn-ami-hvm-2016.09.1.20170119-x86_64-ebs",
  ImageOwnerAlias: "amazon",
  ImageType: "machine",
  Name: "amzn-ami-hvm-2016.09.1.20170119-x86_64-ebs",
  OwnerId: "137112412989",
  Public: true,
  RootDeviceName: "/dev/xvda",
  RootDeviceType: "ebs",
  SriovNetSupport: "simple",
  State: "available",
  VirtualizationType: "hvm"
}
```

## Install

- macos

```bash
$ curl -sL -o ~/bin/alid `curl -s 'http://grl.i-o.sh/youyo/alid?suffix=darwin_amd64'`
$ chmod 755 ~/bin/alid
```

- linux

```bash
$ curl -sL -o ~/bin/alid `curl -s 'http://grl.i-o.sh/youyo/alid?suffix=linux_amd64'`
$ chmod 755 ~/bin/alid
```

## Contribution

1. Fork ([https://github.com/youyo/alid/fork](https://github.com/youyo/alid/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `make lint` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[youyo](https://github.com/youyo)
