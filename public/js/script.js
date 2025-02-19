const configForm = document.getElementById("config-form");
const successAlert = document.getElementById("success-alert");
const failedAlert = document.getElementById("failed-alert");
const threshold = document.getElementById("threshold");
const thresholdUnit = document.getElementById("threshold_unit");
const cronInterval = document.getElementById("cron");
const cronUnit = document.getElementById("cron_unit");

configForm.addEventListener("submit", async (e) => {
  e.preventDefault();

  const tu = thresholdUnit.options[thresholdUnit.selectedIndex].value;
  const cu = cronUnit.options[cronUnit.selectedIndex].value;
  const data = {
    threshold: parseInt(threshold.value),
    threshold_unit: tu,
    cron_interval: parseInt(cronInterval.value),
    cron_unit: cu,
  };

  console.log(data);

  const url = "/config";
  try {
    const res = await fetch(url, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
      credentials: "same-origin",
    });

    if (!res.ok) {
      failedAlert.classList.toggle("d-none");
      setTimeout(() => {
        failedAlert.classList.toggle("d-none");
      }, 2000);
    } else {
      successAlert.classList.toggle("d-none");
      setTimeout(() => {
        successAlert.classList.toggle("d-none");
      }, 2000);
    }
  } catch (error) {
    failedAlert.classList.toggle("d-none");
    setTimeout(() => {
      failedAlert.classList.toggle("d-none");
    }, 2000);
  }
});
