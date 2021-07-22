## Hyperledger Fabric TAG 개발 환경 구축

##### 로컬환경 Vagrant 

##### (가상서버셋팅) WindowsOS
pre virtualbox 와 vagrant 설치 필요  

```
git clone https://github.com/donjilgga/Tag.git
cd Tag\vagrant\node01
vagrant up
```

가상서버 접속  
```
vagrant ssh 
```


##### 1 HLF 네트워크 구축 
```
sudo git clone https://github.com/donjilgga/Tag.git
cd /home/vagrant/Tag/network
sudo chmod +x *.sh
sudo ./build.sh
```


##### 2 HLF 채널생성 & Tag 체인코드 배포 
```
sudo ./javabuild.sh
```

##### 3 HLF JAVA SDK WEB SERVER 실행
```
cd  /home/vagrant/data/

sudo java -jar -Dfile.encoding=UTF-8 Tag-0.0.1-SNAPSHOT.jar
```

##### !! data 폴더에 톰캣이 없을때
윈도우환경으로 다시나가 vagrant reload 한다
```
exit
vagrant reload
```


##### 4-1. HLF 샘플 초기 데이터 저장하기

```
curl --location --request PUT 'http://localhost:28080/product' \
--header 'Content-Type: application/json' \
--data-raw '{"tagid" : "04:5D:45:32:1E:70:81",
"productcode" : "80231172",
"modelname" : "미니 퀼팅 램스킨 롤라 백",
"brand" : "BURBERRY",
"photo" : "https://assets.burberry.com/is/image/Burberryltd/063d3c67ce20919a4005e8bebd77e291aa02d75d.jpg?$BBY_V2_SL_1x1$&wid=1251&hei=1251",
"color" : "BLACK"}'
```

##### 4-2. HLF 샘플 초기 데이터 불러오기
```
curl --location --request GET 'http://localhost:28080/product/0000' \
--header 'Content-Type: application/json' \

// 결과

vagrant@node01:~/Tag/network$ curl --location --request GET 'http://localhost:28080/product/0000' \ --header 'Content-Type: application/json'
{"brand":"ZERO","color":"BLACK","date":"2021-05-28T17:39:21.635334974Z","modelname":"ZERO000","photo":"https://file.mk.co.kr/meet/neds/2017/01/image_readtop_2017_51822_14850815582756240.jpg","productcode":"0000","tagid":"0000"}curl: (6) Could not resolve host:  --header
curl: (6) Could not resolve host: Content-Type

```

##### !! 데이터를 불러오지 못할경우 
톰캣 설정 포트확인 

data\apache-tomcat-8.5.66\conf\server.xml 
파일에서 확인 개발시 자주쓰는 8080포트를 피하기 위해 28080으로  설정하였다
```
<Connector port="28080" protocol="HTTP/1.1"
               connectionTimeout="20000"
               redirectPort="8443" />
```


##### 5. dApp & api server
app 폴더에   
Tag.zip 안드로이드 소스  
fabric-web-server.zip webserver(api server) 소스  
백업함  


##### 6. 블록체인 네트워크 모니터링 스크립트
도커 네트워크 모니터링  
```
cd ~/Tag/network
./monitordocker.sh
```
