#!/usr/bin/env bash
# wait-for-it.sh

set -e

host="$1"
port="$2"
shift 2
cmd="$@"

echo "Aguardando $host:$port ficar disponível..."

until nc -z "$host" "$port"; do
  sleep 1
done

echo "$host:$port está disponível, iniciando comando..."
exec $cmd
