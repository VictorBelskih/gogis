function createChart(totalAreaByFieldType) {
  var obj = Object.values(totalAreaByFieldType);
  var ctx = document.getElementById('myChart').getContext('2d');
  var data = {
      labels: ['Пашня', 'Сенокос', 'Пастбище', 'Залежь'],
      datasets: [{
          label: 'Площадь сельхоз угодий', // Название набора данных
          data: obj,
          backgroundColor: ['#ffbf00', '#c0ff00', '#ff3f00', '#00c0ff']
      }]
  };
  
  var myChart = new Chart(ctx, {
      type: 'pie',
      data: data,
      options: {
          plugins: {
              title: {
                  display: true,
                  text: 'Площадь сельхоз угодий', // Название графика
                  position: 'top'
              },
              legend: {
                  display: true,
                  position: 'bottom' // Позиция легенды
              }
          }
      }
  });
}


function avgOrganicStackedChart(avgOrganic) {
  const labels = avgOrganic.map(item => item.Class);
  const values = avgOrganic.map(item => item.TotalArea);
  
  const ctx = document.getElementById('myPieChart').getContext('2d');
  
  const myPieChart = new Chart(ctx, {
    type: 'bar',
    data: {
      labels: labels,
      datasets: [{
        label: 'Гектар',
        data: values,
        backgroundColor: ['rgb(255, 99, 132)', 'rgb(54, 162, 235)', 'rgb(255, 205, 86)'],
      }]
    },
    options: {
      responsive: true,
      plugins: {
        legend: {
          position: 'bottom',
        },
        title: {
          text: 'Среднее содержание гумуса',
          display: true,
        }
      }
    }
  });
}
function createPieChart(soilData) {
  const labels = soilData.map(item => item.Type + ' ' + item.SubType);
  const values = soilData.map(item => item.Area);
  
  const ctx = document.getElementById('soilChart').getContext('2d');
  
  const myPieChart = new Chart(ctx, {
    type: 'pie',
    data: {
      labels: labels,
      datasets: [{
        data: values,
        backgroundColor: ['rgb(255, 99, 132)', 'rgb(54, 162, 235)', 'rgb(255, 205, 86)', 'rgb(75, 192, 192)', 'rgb(153, 102, 255)', 'rgb(201, 203, 207)', 'rgb(255, 159, 64)', 'rgb(255, 99, 132)', 'rgb(54, 162, 235)', 'rgb(255, 205, 86)'],
      }]
    },
    options: {
      responsive: true,
      plugins: {
        legend: {
          display: false,
          position: 'bottom',
        },
        title: {
          display: true,
          text: 'Преобладающие типы почв'
        }
      }
    }
  });
}
function avgKStackedChart(avgK) {
  const labels = avgK.map(item => item.Class);
  const values = avgK.map(item => item.TotalArea);
  
  const ctx = document.getElementById('avgK').getContext('2d');
  
  const myPieChart = new Chart(ctx, {
    type: 'bar',
    data: {
      labels: labels,
      datasets: [{
        label: 'Гектар',
        data: values,
        backgroundColor: ['rgb(255, 99, 132)', 'rgb(54, 162, 235)', 'rgb(255, 205, 86)'],
      }]
    },
    options: {
      responsive: true,
      plugins: {
        legend: {
          position: 'bottom',
        },
        title: {
          text: 'Среднее содержание калия',
          display: true,
        }
      }
    }
  });
}

function avgPStackedChart(avgP) {
  const labels = avgP.map(item => item.Class);
  const values = avgP.map(item => item.TotalArea);
  
  const ctx = document.getElementById('avgP').getContext('2d');
  
  const myPieChart = new Chart(ctx, {
    type: 'bar',
    data: {
      labels: labels,
      datasets: [{
        label: 'Гектар',
        data: values,
        backgroundColor: ['rgb(255, 99, 132)', 'rgb(54, 162, 235)', 'rgb(255, 205, 86)'],
      }]
    },
    options: {
      responsive: true,
      plugins: {
        legend: {
          position: 'bottom',
        },
        title: {
          text: 'Среднее содержание фосфора',
          display: true,
        }
      }
    }
  });
}
