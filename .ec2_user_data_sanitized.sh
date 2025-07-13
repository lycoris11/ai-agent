#!/bin/bash
yum update -y
yum install -y docker git
sudo yum install python3 -y
sudo yum install python3-pip -y

service docker start
systemctl enable docker
usermod -a -G docker ec2-user

sleep 10

# Clone the repo and build/run
cd /home/ec2-user
git clone https://github.com/lycoris11/ai-agent
cd ai-agent
cat <<EOF > /home/ec2-user/.env
OPENAI_API_KEY=
WEATHER_API_KEY=
HEY_GEN_VIDEO_API_KEY=
ENV=prod
REFRESH_TOKEN=
CLIENT_ID=
CLIENT_SECRET=
EOF
chown ec2-user:ec2-user /home/ec2-user/.env
docker build -t ai-agent:latest .
docker run -d -p 80:8080 --env-file /home/ec2-user/.env ai-agent:latest
#docker run -d -p 127.0.0.1:8080:8080 --env-file /home/ec2-user/.env ai-agent:latest