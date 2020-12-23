FROM postgres:latest

RUN export LC_ALL="ja_JP.UTF-8"
RUN export LC_CTYPE="ja_JP.UTF-8"
RUN localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8
ENV LANG ja_JP.UTF-8
