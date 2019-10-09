source ~/.profile
echo ''
echo 'Checking prerequisite:'
echo ''
echo -n 'Curl:           '
curl --version
echo ''
echo -n 'Node:           '
node --version
echo -n 'npm:            '
npm --version
echo -n 'Docker:         '
docker --version
echo -n 'Docker Compose: '
docker-compose --version
echo -n 'Python:         '
python -V
echo -n 'GoLang          '
go version
