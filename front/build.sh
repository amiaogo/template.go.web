#!/usr/bin/env bash

FINDNAME=$0
while [ -h $FINDNAME ] ; do FINDNAME=`ls -ld $FINDNAME | awk '{print $NF}'` ; done
RUNDIR=`echo $FINDNAME | sed -e 's@/[^/]*$@@'`
unset FINDNAME

# cd to top level agent home
if test -d $RUNDIR; then
  cd $RUNDIR/..
else
  echo 'ERROR'
  exit 1
fi

cd front

npm install

rm -rf ../static/dist
rm -rf ../static/index.html
npm run build
if [ ! -d $(pwd)/../views ]; then
  mkdir $(pwd)/../views
fi
mv ../static/index.html ../views/index.html

echo '前端构建结束'
