---
title: Example Flows
---
:::info
You can apply theses flows multiple times to stay updated, however this will discard all changes you've made.
:::

## Enrollment (2 Stage)

Flow: right-click [here](/static/flows/enrollment-2-stage.json) and save the file.

Sign-up flow for new users, which prompts them for their username, email, password and name. No verification is done. Users are also immediately logged on after this flow.

## Enrollment with email verification

Flow: right-click [here](/static/flows/enrollment-email-verification.json) and save the file.

Same flow as above, with an extra email verification stage.

You'll probably have to adjust the Email stage and set your connection details.

## Two-factor Login

Flow: right-click [here](/static/flows/login-2fa.json) and save the file.

Login flow which follows the default pattern (username/email, then password), but also checks for the user's OTP token, if they have one configured

## Login with conditional Captcha

Flow: right-click [here](/static/flows/login-conditional-captcha.json) and save the file.

Login flow which conditionally shows the users a captcha, based on the reputation of their IP and Username.

By default, the captcha test keys are used. You can get a proper key [here](https://www.google.com/recaptcha/intro/v3.html)

## Recovery with email verification

Flow: right-click [here](/static/flows/recovery-email-verification.json) and save the file.

Recovery flow, the user is sent an email after they've identified themselves. After they click on the link in the email, they are prompted for a new password and immediately logged on.

## User deletion

Flow: right-click [here](/static/flows/unenrollment.json) and save the file.

Flow for users to delete their account,

:::warning
This is done without any warning.
:::
