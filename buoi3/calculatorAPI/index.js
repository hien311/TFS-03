const buttons = document.querySelectorAll(".buttons-item:not(.done)");
const result = document.querySelectorAll(".result")
const equals = document.querySelector(".done")
const ac = document.querySelector(".ac")
const del = document.querySelector(".del")
const power = document.querySelector('.power')
let isTyping = true
let isPower = true;
// css result screen

result.onclick = function(e) {
    console.log(e.target)
}

for (var button of buttons) {
    button.onclick = function(e) {
        if (result[0].innerHTML == "" && [" + "," - "," x "," / ","."].includes(e.target.innerHTML)) {
        }else {
            if (result[0].innerHTML[result[0].innerHTML.length-1] == "." && e.target.innerHTML == ".") {
            }else 
            {
                if (isTyping) {
                    result[0].innerHTML += e.target.innerHTML 
                } else {
                        if ([" + "," - "," x "," / "].includes(e.target.innerHTML)) {
                            result[0].innerHTML = result[1].innerHTML + e.target.innerHTML 
                            result[1].classList.remove("focused")
                            result[0].classList.add("focused")
                            isTyping = true
                        }
                }
            }
        }
    }
}

ac.onclick = function(e) {
    result[0].innerHTML = "";
    result[1].innerHTML = "";
    result[1].classList.remove("focused")
    result[0].classList.add("focused")
    isTyping = true;
}

del.onclick = function(e) {
    if (isTyping) {
        if (result[0].innerHTML[result[0].innerHTML.length - 1 ] == " "){
            result[0].innerHTML = result[0].innerHTML.slice(0,-3)
        } else {
            result[0].innerHTML = result[0].innerHTML.slice(0,-1)
        }
    }
}

power.onclick = function(e) {
    isPower = !isPower;
    const screen = document.querySelector('.screen');
    screen.classList.toggle("display", isPower)
}


equals.onclick = function() {
    if (result[0].innerHTML != "" && result[0].innerHTML[result[0].innerHTML.length-1] != " ") {
        let a = result[0].innerHTML;
        isTyping = false
        // css result screen
        result[0].classList.remove("focused");
        result[1].classList.add("focused")
        // tinh tong?
        fetch('http://localhost:3000/result', {
            method: 'POST',
            body: JSON.stringify(a)
        })
        .then(response => response.json())
        .then(data => {
            console.log(data)
        })
    }
}

