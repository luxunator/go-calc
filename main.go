package main

import "github.com/webview/webview"

func main() {
	debug := false
	window := webview.New(debug)
	defer window.Destroy()

	window.SetSize(400, 530, webview.HintFixed)
	window.SetTitle("Calculator")
	// Injecting HTML/CSS Into Window
	window.Navigate(`data:text/html,<!DOCTYPE html><html><head> <link href="https://fonts.googleapis.com/css2?family=Lato&display=swap" rel="stylesheet"><style> * { margin: 0; }.calc {width: 400px;height: 530px; font-family: 'Lato', sans-serif;background-color: #2D3A44; position: absolute;} .display { width: 325px; height: 75px; background-color: #ffffff; border-radius: 10px; text-align: center; margin: 50px auto 40px auto; } .display h{ height: 75px; padding: 0px 10px; font-size: 40px; display: flex; justify-content: flex-end; align-items: center; } .input { display: flex; flex-wrap: wrap; justify-content: space-between; margin: 0 37.5px 0 42.5px; } .button { color: white; width: 70px; height: 70px; font-size: 40px; margin: 0px 5px 15px 0px; display: flex; border-radius: 10px; justify-content:center; align-items:center; text-decoration: none; background-color: #33A5D7; vertical-align: middle; } .button:hover { background-color: #2980b9; } #zero { width: 150px; }</style></head><body><div class="calc"> <div class="display"> <h id="output"></h> </div> <div class="input"> <a class="button" href="#" class="button" onclick="update('7')">7</a> <a class="button" href="#" class="button" onclick="update('8')">8</a> <a class="button" href="#" class="button" onclick="update('9')">9</a> <a class="button" href="#" class="button" onclick="update('/')">&div;</a> <a class="button" href="#" class="button" onclick="update('4')">4</a> <a class="button" href="#" class="button" onclick="update('5')">5</a> <a class="button" href="#" class="button" onclick="update('6')">6</a> <a class="button" href="#" class="button" onclick="update('*')">&times;</a> <a class="button" href="#" class="button" onclick="update('1')">1</a> <a class="button" href="#" class="button" onclick="update('2')">2</a> <a class="button" href="#" class="button" onclick="update('3')">3</a> <a class="button" href="#" class="button" onclick="update('-')">&minus;</a> <a class="button" href="#" class="button" onclick="update('0')">0</a> <a class="button" href="#" class="button" onclick="clearscr()">c</a> <a class="button" href="#" class="button" onclick="calculate()">&#61;</a> <a class="button" href="#" class="button" onclick="update('%2B')">&#43;</a> </div></div></body></html>`)
	// Injecting Javascript Into Window
	window.Init(`var output="";function update(str){if(output.length<10){output+=str;document.getElementById("output").innerHTML=filter(output)}}function calculate(){try{var answer=Number(Math.round(parseFloat(eval(output))+'e5')+'e-5').toString()}catch(err){var answer=output}if(answer!=undefined){document.getElementById("output").innerHTML=answer;output=answer}}function clearscr(){output="";document.getElementById("output").innerHTML=""}function filter(str){str=str.toString().split("/").join("&div;").split("*").join("&times;").split("-").join("&minus;").split("+").join("&plus;");return str}`)
	
	window.Run()
}