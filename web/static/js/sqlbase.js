async function approveQuery(event) {
    const queryName = event.target.getAttribute("data-query-name");
    console.log("Approving query:", queryName);

    fetch(`/approve?query_name=${queryName}`)
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            console.log("Query approval response:", data);
            if (data.success) { // Check for success
                // Redirect to the approved query
                window.location.href = `/open?query_name=sqlbase:${queryName}:query`;
            } else {
                console.error("Error approving query:", data.error);
            }
        })
        .catch(error => {
            console.error("Error approving query:", error);
        });
}

function searchQueries(event) {
    event.preventDefault();

    const searchValue = document.getElementById("search").value;
    if (searchValue.trim() === "") {
        return;
    }
    fetch(`/api/search?query=${searchValue}`)
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            console.log("Search results:", data);
            const resultsTable = document.getElementById("results");
            const resultsBody = resultsTable.getElementsByTagName("tbody")[0];
            resultsBody.innerHTML = "";
            data.forEach(result => {
                const row = document.createElement("tr");
                const nameCell = document.createElement("td");
                const nameLink = document.createElement("a");
                nameLink.innerText = result.name;
                nameLink.href = result.openUrl;
                nameCell.appendChild(nameLink);
                row.appendChild(nameCell);

                // Description
                const descriptionCell = document.createElement("td");
                descriptionCell.innerText = result.description;
                row.appendChild(descriptionCell);

                // Version
                const versionCell = document.createElement("td");
                versionCell.innerText = result.version;
                row.appendChild(versionCell);


                // Add the status column
                const statusCell = document.createElement("td");
                const statusBadge = document.createElement("span");
                statusBadge.classList.add("badge");
                statusBadge.classList.add(result.status === "pending" ? "bg-warning" : "bg-success");
                statusBadge.innerText = result.status === "pending" ? "Pending" : "Approved";
                statusCell.appendChild(statusBadge);
                row.appendChild(statusCell);

                // Add the "Approve" button for pending queries
                if (result.status === "pending") {
                    const approveButton = document.createElement("button");
                    approveButton.classList.add("badge");
                    approveButton.classList.add("bg-primary"); // Change the color
                    approveButton.style.border = "none";
                    approveButton.style.cursor = "pointer";
                    approveButton.style.marginLeft = "8px"; // Add padding to the left
                    approveButton.style.padding = "4px 8px"; // Add some padding to the button
                    approveButton.style.borderRadius = "4px"; // Add some border radius to make it more button-like
                    approveButton.setAttribute("data-query-name", result.name);
                    approveButton.innerText = "Approve";
                    approveButton.onclick = (event) => approveQuery(event);
                    statusCell.appendChild(approveButton);
                }

                resultsBody.appendChild(row);
            });
            resultsTable.style.display = "table";
        })
        .catch(error => {
            console.error("Error fetching search results:", error);
        });
}

document.getElementById("newQueryForm").addEventListener("submit", async function (event) {
    event.preventDefault();
    const queryName = document.getElementById("query_name").value;
    const queryNameError = document.getElementById("query_name_error");

    if (/[A-Z]/.test(queryName)) {
        queryNameError.style.display = "block";
        return;
    } else {
        queryNameError.style.display = "none";
    }


    const formData = new FormData(event.target);
    const response = await fetch("/sqlbase", {
        method: "POST",
        body: formData
    });
    if (response.ok) {
        const jsonResponse = await response.json();
        if (jsonResponse.message === 'Query updated successfully') {
            // Show the badge and hide it after 3 seconds
            const badge = document.getElementById("querySubmitBadge");
            badge.style.display = "block";
            setTimeout(() => {
                badge.style.display = "none";
            }, 3000);
        } else {
            console.error("Error creating query:", jsonResponse.message);
        }
    } else {
        console.error("Error creating query:", response.statusText);
    }
});