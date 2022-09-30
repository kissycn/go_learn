#!/bin/zsh

echo "What is your favourite OS?"
select name in "Linux" "Windows" "Mac OS" "UNIX" "Android"; do
  case $name in
  "Linux")
      echo "Linux是一个类UNIX操作系统，它开源免费，运行在各种服务器设备和嵌入式设备。"
      break
    ;;
  "Windows")
    echo "Windows是微软开发的个人电脑操作系统，它是闭源收费的。"
    break
    ;;
  esac
done