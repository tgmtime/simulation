# simulation

su: 220 ~ 300 ~ 380 80* 5sn 

elektrik: 700 ~ 1000 ~ 1300  300*  5 sn

gaz: 600 ~ 1000 ~ 1400  400*  5 sn

device_id: 1'den 4'te kadar id sahip 4 adet divece id değerini alır ve string türünde
device_tag:"water", 
device_statu: en son eklenen 3 iot_value değerinin ortalaması 315 büyükse  "Errors" atanacak  küçükse "No Errors" atacanacak
device_location:"water-map"
iot_value: 280 ~ 325 arasında bir değer alır ve string olarak atanacak
iot_value_type: "m³"


device_id: 5'ten 8'e kadar id sahip 4 adet divece id değerini alır ve string türünde
device_tag:"gass", 
device_statu: en son eklenen 3 iot_value değerinin ortalaması 1000'den büyükse  "Errors" atanacak  küçükse "No Errors" atacanacak
device_location:"gass-map"
iot_value: 960 ~ 1015 arasında bir değer alır ve string olarak atanacak
iot_value_type: "m³"


device_id: 9'dan 12'ye kadar id sahip 4 adet divece id değerini alır ve string türünde
device_tag:"gass", 
device_statu: en son eklenen 3 iot_value değerinin ortalaması 1050'den büyükse  "Errors" atanacak  küçükse "No Errors" atacanacak
device_location:"gass-map"
iot_value: 990 ~ 1170 arasında bir değer alır ve string olarak atanacak
iot_value_type: "kWh"

device_tag:"electric"    iot_value: 990 ~ 1170 arasında bir değer alır ve iot_value_type: "kWh" olur

device_tag:"electric"   iot_value: 960 ~ 1015 arasında bir değer alır ve iot_value_type: "m³" olur



{
  "device_id": "string",//1-8 arasında olmalı
  "device_tag": "string", //elektrik,gass,water
  "device_statu": "string", //arızalı olıup olmaması,
  "device_location": "string",
  "iot_value": "string", 
  "iot_value_type": "string",
  "issueId": "i"
}

      type: 'mysql',
      host: '**.**.***.***',
      port: 3306,
      username: 'egitimbu_night',
      database: 'egitimbu_nightwatch',
      password: '*****'


bu şekilde oluşturulmuş table her 5 sn de bir totalde 12 divece belirtilen prosüdürlerde tabloya insert eden kodu oluşturur musun golang de 

Prosedürler:{
device_id: 1'den 4'te kadar id sahip 4 adet divece id değerini alır ve string türünde
device_tag:"water", 
device_statu: en son eklenen 3 iot_value değerinin ortalaması 315 büyükse  "Errors" atanacak  küçükse "No Errors" atacanacak
device_location:"water-map"
iot_value: 280 ~ 325 arasında random bir değer alır ve string olarak atanacak
iot_value_type: "m³"


device_id: 5'ten 8'e kadar id sahip 4 adet divece id değerini alır ve string türünde
device_tag:"gass", 
device_statu: en son eklenen 3 iot_value değerinin ortalaması 1000'den büyükse  "Errors" atanacak  küçükse "No Errors" atacanacak
device_location:"gass-map"
iot_value: 960 ~ 1015 arasında random bir değer alır ve string olarak atanacak
iot_value_type: "m³"


device_id: 9'dan 12'ye kadar id sahip 4 adet divece id değerini alır ve string türünde
device_tag:"electric", 
device_statu: en son eklenen 3 iot_value değerinin ortalaması 1050'den büyükse  "Errors" atanacak  küçükse "No Errors" atacanacak
device_location:"electric-map"
iot_value: 990 ~ 1170 arasında random bir değer alır ve string olarak atanacak
iot_value_type: "kWh" }


db_connection:{
    type: 'mysql',
      host: '**.**.***.***',
      port: 3306,
      username: 'egitimbu_night',
      database: 'egitimbu_nightwatch',
      password: '*****'
}

CREATE TABLE `iot_entity` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `deleted_at` datetime(6) DEFAULT NULL,
  `device_id` varchar(255) NOT NULL,
  `device_tag` varchar(255) NOT NULL,
  `device_statu` varchar(255) NOT NULL,
  `device_location` varchar(255) NOT NULL,
  `iot_value` varchar(255) NOT NULL,
  `iot_value_type` varchar(255) NOT NULL,
  `issueId` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_b54a14e9ba193ae688324d3443` (`deleted_at`),
  KEY `FK_26c1ef738fefb4f9c959a01c684` (`issueId`),
  CONSTRAINT `FK_26c1ef738fefb4f9c959a01c684` FOREIGN KEY (`issueId`) REFERENCES `issue` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;




 Endüstri 4.0, Yapay zeka, bulut sistemler, big data gibi teknolojilerin kullanıldığı veriler
ile oluşturulan ve kullanılan bir devrimdir. Endüstri 5.0 ise insanların yaşam kalitesini 
arttırmak için süper toplumların oluşmasına katkıda bulunan bir ileri teknolojik devrimdir.


 Endüstri 4.0 temel olarak bir işin nasıl daha iyi yapıldığıyla ilgilenirken, Endüstri 5.0 ise yapılan 
işin zaman – insan ilişkisi içerisinde nasıl en verimli hale getirileceği ile ilgilenir. Endüstri 4.0 
otomatikleşen makinelerin kullanımını yaygınlaştırırken ve bunu ilke olarak edinirken, Endüstri 5.0 ise 
insan – makine arasında hızlıca öğrenen bir yapı kurmayı amaç edinir ve bu sayede problemin çözümü için 
en optimum durumu arar.



1- case
