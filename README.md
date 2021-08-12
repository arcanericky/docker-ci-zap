# Windows Docker Layer Zapper

A utility to obliterate the `C:\ProgramData\docker` directory when nothing else will work.

[![Build](https://github.com/arcanericky/docker-ci-zap/actions/workflows/builder.yml/badge.svg)](https://github.com/arcanericky/docker-ci-zap/actions/workflows/builder.yml)
![GitHub License](https://img.shields.io/github/license/arcanericky/dockercizap)

## WARNING

Before using this utility you should read a couple of posts on the dangers of using it, including _the possibility that you might "break the operating system"_. Instead, consider using the GUI and navigating to _Settings_, then _Reset_. Good luck.

- [Dangerous Utility][danger]
- [Clean up after yourself Docker][cleanup]
- [Disk space full running your Sitecore Docker instances][diskfull]

## Dependencies

This package relies on the [hcsshim package][hcsshim] for the heavy lifting.
It's the Go interface for using the [Windows Host Compute Service][win-hcs] (HCS) to
launch and manage Windows Containers.

## Quick Start

The original documentation for using this utility assumes the desire to obliterate everything in `C:\ProgramData\docker` and start over. That's probably the safest route because metadata for `image\` and `windowsfilter\` is scattered throughout this directory structure.

If you get errors (exceptions) when executing this utility, the cleanest technique is to uninstall Docker, reboot, then run the utility with administrative privileges. I've also found that stopping the `Docker Engine` and `Docker Desktop Service` will usually eliminate these exceptions which are usually caused by the lower level hcshim package and not preventable in the code.

Execute:

```cmd
.\docker-ci-zap.exe -folder "C:\ProgramData\docker"
```

## Why?

If you've used [Docker for Windows][docker-windows] with [Windows containers][windows-containers] you've probably realized that Windows images are huge and the layers are by default packed into `C:\ProgramData\Docker\image` and `C:\ProgramData\Docker\windowsfilter` directories. Eventually you run low on disk space, notice those directories seem to be responsible, then try to remove it only to find you can't do this. Next, you attempt removal with administrator privileges only to fail yet again. This leads to ultimate sadness and in desperation you search for relief. If you're lucky you find the executable for the [original `docker-ci-zap` project][docker-ci-zap] from which this repository was forked. Embedded in that repository is an executable that will save you.

Microsoft's [Container Storage Overview][container-overview] post will start you down the rabbit hole of how all this works and also notably references how to change where these layers are stored using the `docker-root` configuration.

The [original project][docker-ci-zap] hasn't been updated in years and there are no formal releases or testing. It's also tough to tell which version the [hcsshim package][hcsshim] was used for the build and the executable in the original repository shows it was lasted updated 5 years ago.

This project forks the original, adding more formal build and test steps along with a more formal release using versions. It's an attempt to make this executable more accessible and more safe. Because no one really enjoys downloading and running some random executable packed into a random repository.

[cleanup]: https://freddysblog.com/2018/12/11/clean-up-after-yourself-docker-your-mom-isnt-here/
[container-overview]: https://docs.microsoft.com/en-us/virtualization/windowscontainers/manage-containers/container-storage
[danger]: https://github.com/moby/moby/issues/26873#issuecomment-249338936
[diskfull]: https://visionsincode.com/2021/02/14/disk-space-full-running-your-sitecore-docker-instances/
[docker-ci-zap]: https://github.com/moby/docker-ci-zap
[docker-windows]: https://docs.docker.com/docker-for-windows/install/
[hcsshim]: https://github.com/Microsoft/hcsshim
[win-hcs]: https://docs.microsoft.com/en-us/virtualization/api/hcs/overview
[windows-containers]: https://www.docker.com/products/windows-containers