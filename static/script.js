
const options = { style: 'currency', currency: 'BRL', minimumFractionDigits: 2, maximumFractionDigits: 3 }
const formatNumber = new Intl.NumberFormat('pt-BR', options)
var map = new Map()

for (let i = 0; i <= 200; i++) {
    let element = document.getElementById(i)
    if(element){
        console.log(element)
        map.set(i, true)
        continue
    } 
    map.set(i, false)
}


addEventListener("DOMContentLoaded", () => {
    valores = document.querySelectorAll(".real")
    valores.forEach(valor => {
        console.log()
        valor.textContent = formatNumber.format(valor.textContent)
    });
})

function scrollToId() {
    let Id = document.getElementById('searchPlu').value

    let element = document.getElementById(Id)

    if (element) {
        element.scrollIntoView({ behavior: 'smooth' });
    } else {
        alert("Produto não cadastrado")
    }
}

function scrollToIdByValue(value) {
    let element = document.getElementById(value)

    if (element) {
        element.scrollIntoView({ behavior: 'smooth' });
    } else {
        console.log("Produto não cadastrado")
    }
}

function excluir(num) {
    resposta = confirm(`Certeza que deseja deletar o produto ${num}`)
    if (!resposta) {
        return
    }
    bool = false

    linha = document.querySelector(`#row-${num}`)
    console.log(linha)
    linha.classList.add("fade-out")
    {
    plu = parseInt(num)
    let dados = {
        Plu: plu
    };

    let json = JSON.stringify(dados);

    fetch('/excluir', { // Endpoint para exclusão
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: json,
    })
        .then(response => {
            if (!response.ok) {
                return response.text().then(text => { throw new Error(text) });
            }
            return response.json();
        })
        .then(data => {
            console.log('Excluído:', data);
            // Recarregar a página após a exclusão bem-sucedida
            location.reload();
        })
        .catch((error) => {
            try {
                const errorJson = JSON.parse(error.message);
                console.error('Erro ao excluir:', errorJson.message);
            } catch (e) {
                console.error('Erro ao excluir:', error.message);
            }
        });
    }
}

function editar(num) {
    console.log(num)
    inputName = ".input-plu-" + num
    inputs = document.querySelectorAll(inputName)
    for (let i = 0; i < inputs.length; i++) {
        let inp = inputs[i];
        if (inp.style.display == "none") {
            inp.style.display = "inline"
            continue
        }
        inp.style.display = "none"
    }
}

function exibeNew(){
    element = document.getElementById("newPlu")
    if (!element) {
        return
    }
    console.log(element.style.display)
    if (element.style.display == "none") {
        element.style.display = ""
        return
    }
    element.style.display = "none"
}

function enviarDados(dados) {
    let json = JSON.stringify(dados);
    console.log(json)
    fetch('/editar', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: json,
    })
        .then(response => {
            if (!response.ok) {
                return response.text().then(text => { throw new Error(text) });
            }
            return response.json();
        })
        .then(data => {
            console.log('Sucesso:', data);
            location.reload();
        })
        .catch((error) => {
            try {
                const errorJson = JSON.parse(error.message);
                console.error('Erro:', errorJson.message);
            } catch (e) {
                console.error('Erro:', error.message);
            }
        });
}

function coletarDados(plu) {
    let descricao = document.querySelector(`input[name="descricao-${plu}"]`).value.trim();
    descricao = descricao.toUpperCase()
    let preco = parseFloat(document.querySelector(`input[name="preco-${plu}"]`).value.trim());
    let vendaTexto = document.querySelector(`span.input-plu-${plu}`).textContent.trim();
    let venda = vendaTexto === 'Unitario' ? 1 : 0;
    let validade = parseInt(document.querySelector(`input[name="validade-${plu}"]`).value.trim());

    bolean = descricao == "" || isNaN(preco) || vendaTexto == "" || isNaN(validade)

    if (bolean) {
        console.log("erro cadastro")
        alert("erro cadastro")
        return
    }
    plu = parseInt(plu)


    let dados = {
        Plu: plu,
        Descricao: descricao,
        Preco: preco,
        Venda: venda,
        Validade: validade
    };


    console.log(dados)
    return enviarDados(dados);
}

function coletarTodosDados() {
    let linhas = document.querySelectorAll('tbody tr');
    let dados = [];

    linhas.forEach(linha => {
        let plu = linha.querySelector('td[id]').textContent.trim();
        let dadosPLU = coletarDados(plu);
        dados.push(dadosPLU);
    });

    return enviarDados(dados);
}

function novo() {
    let pluHtml = document.getElementById("plu-new")
    if (!pluHtml) {
        return false
    }

    let descricao = (document.getElementById("descricao-new").value).toUpperCase()
    let preco = parseFloat(document.getElementById("preco-new").value)
    let vendaTexto = document.getElementById("venda-new").value
    let venda = vendaTexto === 'Unitario' ? 1 : 0;
    let validade = parseInt(document.getElementById("validade-new").value)

    console.log(pluHtml.value)
    let pluInvalido = isNaN(pluHtml.value) || descricao == "" || isNaN(preco) || vendaTexto == "" || isNaN(validade) 
    
    if (pluInvalido) {
        alert("Novo cadastro invalido")
        return false
    }

    let plu = parseInt(pluHtml.value) 
    if(map.get(plu)) {
        alert("Plu já cadastrado")
        scrollToIdByValue(plu)
        return false
    }

    let dados = {
        Plu: plu,
        Descricao: descricao,
        Preco: preco,
        Venda: venda,
        Validade: validade
    };

    return enviarDadosNovoPlu(dados)
}


function enviarDadosNovoPlu(dados) {
    let json = JSON.stringify(dados);
    console.log(json)
    fetch('/novo', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: json,
    })
        .then(response => {
            if (!response.ok) {
                return response.text().then(text => { throw new Error(text) });
            }
            return response.json();
        })
        .then(data => {
            console.log('Sucesso:', data);
            location.reload();
        })
        .catch((error) => {
            try {
                const errorJson = JSON.parse(error.message);
                console.error('Erro:', errorJson.message);
            } catch (e) {
                console.error('Erro:', error.message);
            }
        });
}

async function startImport(event) {
    reiniciarCronometro()
    iniciarCronometro()
    event.preventDefault(); // Prevent default form submission
    const form = document.getElementById('importForm');
    const campos = form.querySelectorAll('input[required]');

    let todosPreenchidos = true;
    campos.forEach(campo => {
        if (!campo.value.trim()) {
            todosPreenchidos = false;
        }
    });

    if (!todosPreenchidos) {
        const aviso = 'Por favor, preencha todos os campos obrigatórios.';
        alert(aviso)
        return;
    }

    var progressElement = document.getElementById("progress");
    var statusElement = document.getElementById("status");
    var ellipsisElement = document.getElementById("ellipsis");
    var progressBar = document.querySelector(".progress-bar");

    // Reset progress bar and status
    progressElement.style.width = '0%';
    progressElement.textContent = '0%';
    statusElement.textContent = 'Importando...';
    progressBar.style.display = 'block';

    var ellipsisCounter = 0;
    var ellipsisInterval = setInterval(function () {
        ellipsisElement.textContent = '.'.repeat((ellipsisCounter % 4) + 1);
        ellipsisCounter++;
    }, 500); // Change ellipsis every 500ms

    var pulsateInterval = setInterval(function () {
        statusElement.style.fontSize = '1.01em';
        setTimeout(function () {
            statusElement.style.fontSize = '1em';
        }, 250);
    }, 500); // Pulsate every 500ms

    var formData = new FormData(document.getElementById("importForm"));
    

    var request = new XMLHttpRequest();
    request.open("POST", "/importarteste", true);
    request.setRequestHeader("Accept", "text/event-stream");

    request.onreadystatechange = function () {
        if (request.readyState === 3 || request.readyState === 4) { 
            var lines = request.responseText.split("\n\n");
            lines.forEach(function (line) {
                if (line.startsWith("data: ")) {
                    var data = line.replace("data: ", "").trim();
                    if (data) {
                        try {
                            var jsonData = JSON.parse(data);
                            if (jsonData.hasOwnProperty("progress")) {
                                var progressValue = jsonData.progress;
                                progressElement.style.width = progressValue + '%';
                                progressElement.textContent = Math.round(progressValue) + '%';
                            }
                            if (jsonData.hasOwnProperty("complete")) {
                                clearInterval(ellipsisInterval); // Stop ellipsis animation
                                clearInterval(pulsateInterval); // Stop pulsate animation
                                ellipsisElement.textContent = ''; // Clear ellipsis
                                statusElement.textContent = 'Finalizado';
                                reiniciarCronometro()
                            }
                        } catch (e) {
                            // Handle non-JSON data
                            console.error("Error parsing JSON data:", e);
                        }
                    }
                }
            });
        }
        if (request.readyState === 4 && request.status !== 200) {
            statusElement.textContent = 'Erro ao completar a tarefa.';
            reiniciarCronometro()
            clearInterval(ellipsisInterval); // Stop ellipsis animation
            clearInterval(pulsateInterval); // Stop pulsate animation
            ellipsisElement.textContent = ''; // Clear ellipsis
        }
    };
    request.send(formData);
}

document.getElementById("importForm").addEventListener("submit", startImport);

function formatarData(eventDate){
    var parts = eventDate.split(" "); 
    // Divide a string na primeira ocorrência de espaço
    var complete_time = parts[1];
    var data_particionada = parts[0].split("-");
    var ano = data_particionada[0]
    var mes = data_particionada[1]
    var dia = data_particionada[2]

    var data = `${dia}/${mes}/${ano}`
    
    var time = complete_time.split(".")[0];
    // Agora parts[0] contém "2024-06-25"
    var dataFormatada =`Dia: ${data} às: ${time}`;
    
    return dataFormatada
}

let tempoSegundos = 0;
let cronometroAtivo = false;
let intervalo;

function formatarTempo(segundos) {
    const horas = String(Math.floor(segundos / 3600)).padStart(2, '0');
    const minutos = String(Math.floor((segundos % 3600) / 60)).padStart(2, '0');
    const seg = String(segundos % 60).padStart(2, '0');
    return `${horas}:${minutos}:${seg}`;
}

function atualizarDisplay() {
    document.getElementById('tempo').innerText = formatarTempo(tempoSegundos);
}

async function iniciarCronometro() {
    if (cronometroAtivo) return;
    cronometroAtivo = true;
    intervalo = setInterval(() => {
        tempoSegundos++;
        atualizarDisplay();
    }, 1000);
}

function pausarCronometro() {
    cronometroAtivo = false;
    clearInterval(intervalo);
}

function reiniciarCronometro() {
    pausarCronometro();
    tempoSegundos = 0;
    atualizarDisplay();
}
