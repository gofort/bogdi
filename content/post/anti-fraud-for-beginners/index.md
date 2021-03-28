---
title: "Anti-fraud for Beginners"
description: "3 Simple steps to stop money loss"
date: "2021-03-28"
slug: "anti-fraud-beginners"
tags:
  - security
  - note
---

Some days ago we have faced a fraud. One of our new clients have been creating one new account every week to use our product for free. So we needed mechanisms to prevent this behaviour. IP addresses which this client used were different each time, so 'block by IP' doesn't work in this case. Bank cards were also different each time. At the same time, we needed to fix this as soon as possible to prevent money loss, so we came to the following things: 

* **Enable 3D Secure**. In this case bank is responsible for transaction, without 3D secure you are responsible for transaction, so fraudsters don't use 3D secure for their cards.
* **Restrict registration using disposable emails**. Fraudsters often use disposable emails, because it is easier to automate new email address creation and nobody will block them, so just restrict registration from their domains. See the list of domains [here](https://github.com/disposable/disposable/blob/master/blacklist.txt).
* **Enable reCaptcha**. It is much harder to automate registration in your platform if you have captcha and it makes process of registration less attractive to fraudsters. It is better to use [reCaptcha](https://www.google.com/recaptcha/about/), because it automatically detects 'bad' users, so it won't affect user flow for everybody else.
