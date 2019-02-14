#!/bin/bash
#echo 'copy wars'
#cp /wars/membermanager*.war /opt/tomcat/webapps/membermanager.war
#cp /wars/breservices*.war /opt/tomcat/webapps/breservices.war
#cp /wars/bre-1*.war /opt/tomcat/webapps/bre.war
#cp /wars/maintenance*.war /opt/tomcat/webapps/maintenance.war
#cp /wars/XStatic*.war /opt/tomcat/webapps/XStatic.war
#cp /wars/axnbusiness*.war /opt/tomcat/webapps/b2b.war
#cp /wars/printservice*.war /opt/tomcat/webapps/printservice.war
#cp /wars/bre-client*.war /opt/tomcat/webapps/breclient.war
 
echo 'touch tomcat catalina.out'
#touch /opt/tomcat/logs/catalina.out
#tail -f /opt/tomcat/logs/catalina.out &
 
echo 'tomcat start'
/opt/tomcat/bin/startup.sh

touch /opt/tomcat/logs/catalina.out
tail -f /opt/tomcat/logs/catalina.out

#echo 'copy hsql'
#cp -fr /wars/hsqldb /
 
#echo 'db start'
#java -cp /opt/tomcat/lib/hsqldb.jar org.hsqldb.Server -database.0 file:/hsqldb/axil -dbname.0 axil
