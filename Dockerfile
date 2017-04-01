# Copyright 2017 Joan Llopis. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.
FROM alpine:3.5
LABEL maintainer="Joan Llopis <jllopis@acb.es>" \
      version=v0.1.0 \
	  description="ZFS backup server that manage jobs and tasks to perform local or remote backups"

ENV ZBS_VERSION v0.0.1
ADD  _build/zbs_${ZBS_VERSION}_srv_linux_amd64.bin /zbs/zbs.bin

WORKDIR /zbs
EXPOSE 8000

ENTRYPOINT ["./zbs.bin"]
