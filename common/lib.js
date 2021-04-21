function generateTable(data) {
    if (data.length === 0)
        return null;

    let headObject = data[0];

    let table = document.createElement("table");
    let headRow = document.createElement("tr");

    table.appendChild(headRow);

    for (let name in headObject) {
        let th = document.createElement("th")
        th.innerText = name;
        headRow.appendChild(th);
    }

    for (let row of data) {
        let tr = document.createElement("tr");
        for (let field in row) {
            let value = row[field];
            let td = document.createElement("td");

            if (value !== null && value !== undefined) {
                td.innerText = value;
            } else {
                td.innerHTML = `<div class="null">null</div>`
            }

            tr.appendChild(td);
        }
        table.appendChild(tr);
    }

    return table;
}

function generateInputForm(data) {
    if (data.length === 0)
        return null;

    let form = document.createElement("form");

    for (let name in data) {
        let label = document.createElement("div")
        label.innerText = name;

        form.appendChild(label);

        let input = document.createElement("input");
        if (typeof data[name] === "number") {
            input.type = "number";
        }
        if (typeof data[name] === "string") {
            input.type = "text";
        }
        if (typeof data[name] === "boolean") {
            input.type = "checkbox";
        }
        input.name = name;
        form.appendChild(input);
    }

    return form;
}

function clear(node) {
    while (node.hasChildNodes()) {
        clear(node.firstChild)
    }
    node.parentNode.removeChild(node);
}

function clearChildren(node) {
    while (node.hasChildNodes()) {
        clear(node.firstChild)
    }
}

function getFormValues(form) {
    let result = {};
    for (let child of form.children) {
        if (child.tagName !== "INPUT")
            continue;
        if (child.value === "") {
            result[child.name] = null;
            continue
        }

        if (child.type === "text") {
            result[child.name] = child.value;
        }

        if (child.type === "number") {
            let number = parseFloat(child.value);

            if (isNaN(number)) {
                result[child.name] = child.value;
            } else {
                result[child.name] = number;
            }
        }

        if (child.type === "checkbox") {
            result[child.name] = child.checked;
        }
    }
    return result;
}