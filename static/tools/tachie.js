const input = document.getElementById("input")
const scale = document.getElementById("scale")
const drog = document.getElementById("drog")
const status = document.getElementById("status")
scale.value = 1
// picture.draggable = true
// picture.ondragover = function (ev) {
//     console.log("drag", ev)
//     // ev.dataTransfer.setData("Text",ev.target.id);
// }
const offset = {
    x: 0, y: 0
}
const tmp = {
    position: "",
    size: ""
}
drog.prePos = {x: 0, y: 0}
drog.onmousedown = (e) => {
    drog.move = true
    drog.pre = {
        x: e.screenX,
        y: e.screenY
    }
}
drog.onmousemove = (e) => {
    if (drog.move) {
        // console.log("now", (picture.prePos.x + e.screenX - picture.pre.x + "px ") + (picture.prePos.y + e.screenY - picture.pre.y + "px"))
        // picture.style.left = picture.prePos.x + e.screenX - picture.pre.x + "px"
        // picture.style.top = picture.prePos.y + e.screenY - picture.pre.y + "px"
        tmp.position = (offset.x + drog.prePos.x + e.screenX - drog.pre.x + "px ") +
            (offset.y + drog.prePos.y + e.screenY - drog.pre.y + "px")
        block.style.backgroundPosition = tmp.position
    }
}
drog.onmouseup = (e) => {
    drog.move = false
    drog.prePos = {
        x: drog.prePos.x + e.screenX - drog.pre.x,
        y: drog.prePos.y + e.screenY - drog.pre.y
    }
}

const block = document.getElementById("block")

const tachie = {
    url: "",
    width: 0,
    height: 0,
    step: 0,
}

function handleInput() {
    status.innerHTML = "head"
    tachie.url = input.value
    const picture = document.createElement("img")
    picture.src = input.value
    tachie.width = picture.width
    tachie.height = picture.height

    block.style.width = "120px"
    block.style.height = "120px"
    block.style.backgroundImage = "url("+input.value+")"
    offset.x = -picture.width/2
    offset.y = -picture.height/2
    block.style.backgroundPosition =
        (offset.x+ "px ") +
        (offset.y+ "px")
}

var setScale = false
function handleScale() {
    setScale = true
}
function handleScaleOver() {
    setScale = false
}
setInterval(() => {
    if (setScale) {
        // picture.style.transform = "scale("+scale.value+")"

        tmp.size =
            (parseInt(tachie.width * scale.value) + "px ") +
            (parseInt(tachie.height * scale.value)+ "px")
        block.style.backgroundSize = tmp.size
    }
}, 40)
const funs = [
    () => {
        tachie.head_transform = {
            position: tmp.position,
            size: tmp.size
        }
        console.log(tachie)
        status.innerHTML = "big"
        block.style.width = "100%"
        block.style.height = "100%"
        offset.x = 0
        offset.y = 0
        block.style.backgroundPosition =
            (offset.x+ "px ") +
            (offset.y+ "px")
        drog.prePos = {x: 0, y: 0}
        return true
    },
    () => {
        tachie.big_transform = {
            position: tmp.position,
            size: tmp.size
        }
        console.log(tachie)
        status.innerHTML = "battle"
        block.style.width = "300px"
        block.style.height = "80px"
        offset.x = -tachie.width/2
        offset.y = -tachie.height/2
        block.style.backgroundPosition =
            (offset.x+ "px ") +
            (offset.y+ "px")
        drog.prePos = {x: 0, y: 0}
        return true
    },
    () => {
        tachie.battle_transform = {
            position: tmp.position,
            size: tmp.size
        }
        console.log(tachie)
        status.innerHTML = "card"
        block.style.width = "200px"
        block.style.height = "260px"
        offset.x = -tachie.width/2
        offset.y = -tachie.height/2
        block.style.backgroundPosition =
            (offset.x+ "px ") +
            (offset.y+ "px")
        drog.prePos = {x: 0, y: 0}
        return true
    },
    () => {
        tachie.card_transform = {
            position: tmp.position,
            size: tmp.size
        }
        console.log(tachie)
        console.log("over!")
        return false
    }
]
function next() {
    if (true) {
        if (funs[tachie.step]()) {
            tachie.step ++
        }
    }
}