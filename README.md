## Introduction:

* This is a Web Application for checking the validation of Iranian national code. 
* Gin Web Framework is used in this program. You can get information about it by using the [this link](https://github.com/gin-gonic/gin).
* The national code is as input, its validation is checked. In the output of the program, the result of the check and also if the code is valid, the name of the city where the ID card was issued is displayed.

## Check the validation of the national code :

The steps to check the validity of the national code are described in [this link](http://aliarash.com/article/codemeli/codemeli.htm) but it is on Persian, so I translated into English:

The national code is a 10-digit number. The first three digits from the left indicate the city code of the place of issuance of the identity card. The next six digits of the code are unique to the person holding the identity card in the city where it was issued, and the last digit is a control digit that is obtained from the nine digits on the left. To check the control digit, it is enough to recalculate the control digit from the nine digits on the left.Note that the first digit and the second digit to the left of the national code may be zero.

1. To calculate the control digit using other digits, we multiply each digit by its position and add the result together.
2. Divide the sum obtained from step one by 11.
3. If the residual of the division is less than two, the control digit must be equal to the residual, otherwise the control digit must be equal to the result of (11-residual).

* For example:

Input: /?id=0049557025 (You can use the Postman to get the result)

Output: 
{
    "City": "Tehran-Center",
    "National-ID": "0049557025",
    "Result": "National-ID is Valid"
}




