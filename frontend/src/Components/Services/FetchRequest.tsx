export async function fetchRequestNoBearer(url = "", data = {}, method = "POST") {
  const response = await fetch(url, {
    method: method,
    mode: "cors",
    cache: "no-cache",
    credentials: "same-origin",
    headers: {
      "Content-Type": "application/json"
    },
    redirect: "follow",
    referrerPolicy: "no-referrer",
    body: JSON.stringify(data)
  });

  if (response.status !== 200) {
    return { status: response.status, errorText: response.statusText };
  }

  return response.json(); // parses JSON response into native JavaScript objects
}

export async function fetchRequest(url = "", data = {}, method = "POST") {
  const response = await fetch(url, {
    method: method,
    mode: "cors",
    cache: "no-cache",
    credentials: "same-origin",
    headers: {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    redirect: "follow",
    referrerPolicy: "same-origin",
    body: method == "GET" ? null : JSON.stringify(data)
  });

  if (response.status !== 200) {
    return { status: response.status };
  }

  return response.json(); // parses JSON response into native JavaScript objects
}
