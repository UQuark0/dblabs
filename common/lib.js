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

            if (value) {
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