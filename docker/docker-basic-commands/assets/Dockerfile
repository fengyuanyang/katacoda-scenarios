FROM fengyuanyang/jdk6-mvn3

ENV MAVEN_OPTS=-Xmx2048m
ENV NODE_OPTIONS=--max_old_space_size=2048
ENV JAVA_OPTS="-Xms512m -Xmx2048m -XX:PermSize=256m -XX:MaxPermSize=1024m"
ENV CATALINA_OPTS="-ea -Daxn.work.root=/BRE/axn/work -Daxn.env.override=/BRE/axn/work/appenvconf/AXNEnvironment.properties -Daxn.testing.integration.persistence-mode=HSQLDB -Dlogback.configurationFile=/BRE/axn/work/config/axnlog4jconfig.xml"

#COPY ./tomcatConf/server.xml /opt/tomcat/conf/server.xml
#COPY ./tomcatJars/*.jar /opt/tomcat/lib/

# enable debuging in tomcat
ENV JPDA_ADDRESS=5005

EXPOSE 8090
EXPOSE 9001
EXPOSE 5005
