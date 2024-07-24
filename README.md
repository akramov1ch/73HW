# Uyga vazifa: RabbitMQ Topic Exchange orqali Xabar Yuborish va Qabul Qilish (Boshqaruv Tizimi)

## Maqsad
Ushbu vazifaning maqsadi `RabbitMQ` ning `topic exchange` xususiyati va maxsus belgilar (*, #) bilan ishlashni o'rganish va uni murakkab tizimlarda qanday qo'llashni tushunishdir.

## Talablar
1. **REST API (Producer) yaratish**: 
    - Yangi xabarlarni yaratish va o'qish uchun endpointlarga ega `REST API` ni amalga oshiring.
    - Yangi xabar yaratilganda, xabar tafsilotlarini `RabbitMQ` `topic exchange` ga yuboring. Har bir xabar tegishli `routing kaliti` bilan jo'natiladi (masalan, "`report.created`", "`report.updated`").

2. **Worker (Consumer) yaratish:**: 
    - Har bir turdagi xabarlarni qayta ishlash uchun alohida consumerlar yarating.
    - Consumer `RabbitMQ` dan mos xabarlarni olishi va qayta ishlashi kerak, masalan:
        - Xabar tafsilotlarini `MongoDB` ga saqlash.
        - Xabar holatini yangilash.

## Qo'shimcha Talablar
- Topic Exchange bilan Ishlash:
    - Xabarlarni yuborishda va qabul qilishda routing kalitlarida maxsus belgilar (masalan, "*", "#") bilan ishlang.
    - Turli routing kalitlaridan foydalanib xabarlarni qabul qiluvchi consumerlar yarating. Misol uchun:
        - `report.*` routing kalitidan foydalanadigan consumer: har qanday report bilan bog'liq xabarlarni qabul qiladi.
        - `*.updated` routing kalitidan foydalanadigan consumer: har qanday yangilangan xabarlarni qabul qiladi.
        - `report.#` routing kalitidan foydalanadigan consumer: barcha report bilan bog'liq xabarlarni qabul qiladi, ichki mavzulardan qat'i nazar.
