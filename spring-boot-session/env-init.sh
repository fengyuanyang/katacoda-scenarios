PROJECT=spring-session	  # The name of the folder within the code samples repo to copy
UI_PATH=/root/code 	  # This should match your index.json key

git clone https://github.com/spring-projects/spring-session.git
git checkout 2.1.4.RELEASE
#https://github.com/spring-projects/spring-session.git
cd ${UI_PATH} && cp -R /root/${PROJECT}/* ./
clear # To clean up Katacoda terminal noise
~/.launch.sh
