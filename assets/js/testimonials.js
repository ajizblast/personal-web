// class Testimonial {
//     #quotes = ""
//     #images = ""

//     constructor(quotes,images){
//         this.#quotes = quotes,
//         this.#images = images
//     }

//     get quotes(){
//         return this.#quotes
//     }

//     get images(){
//         return this.#images
//     }

//     get forHTML(){
//         return `<div class="testi-card">
//         <img src="${this.images}" alt="testi">
//         <div class="testi-desc">
//             <p class="quotes">"${this.quotes}"</p>
//             <p class="author">- ${this.author}</p>
//         </div>
//     </div>`
//     }
// }

// class authorTestimonial extends Testimonial {
//     #author = ""
//     constructor(author,quotes,images){
//         super(quotes,images)
//         this.#author = author
//     }

//     get author(){
//         return this.#author
//     }
// }

// class companyTestimonial extends Testimonial{
//     #company
//     constructor(company,quotes,images){
//         super(quotes,images)
//         this.#company = company
//     }

//     get author() {
//         return this.#company + " Company"
//     }
// }

// const testi1 = new authorTestimonial("Martha Cristina Tiahahu", "Perang bukan satu-satunya menuju merdeka","./assets/images/cristina.jpg")
// const testi2 = new authorTestimonial("Jendral Sudirman", "Mantap sekali ,aku yang pembuat facebook kalah","./assets/images/jendral-sudirman.jpg")
// const testi3 = new companyTestimonial("Bung Karno", "Masnya ini jago UI/UX, jago slicing juga, database juga mantap. paket lengkap ya mas","./assets/images/sukarno.jpg")
// const testi4 = new authorTestimonial("Bung Tomo", "Sepertinya pembuat website ini ingin ku recruit ke perusahaanku","./assets/images/bung-tomo.jpg")
// const testi5 = new authorTestimonial("R.A. Kartini", "Pembuat wbsite andalanku dari dulu hingga nanti","./assets/images/kartini.pg.jpg")
// const testi6 = new authorTestimonial("Kapiten Pattimura", "selalu ada plus minus, tapi sama mas aji plusnya bisa membuat Indonesia merdeka lagi","./assets/images/pattimura.png")

// let data = [testi1,testi2,testi3,testi4,testi5,testi6]
// let testimonialforHTML = ""

// for(let i = 0; i < data.length; i++){
//     testimonialforHTML += data[i].forHTML
// }

// document.getElementById("testimonials").innerHTML = testimonialforHTML

const testimonialData = [
  {
    name: "Martha Cristina Tiahahu",
    quote: "Perang bukan satu-satunya menuju merdeka",
    image: "./assets/images/cristina.jpg",
    rating: 2,
  },
  {
    name: "Jendral Sudirman",
    quote: "Mantap sekali ,aku yang pembuat facebook kalah",
    image: "./assets/images/jendral-sudirman.jpg",
    rating: 5,
  },
  {
    name: "Bung Karno",
    quote:
      "Masnya ini jago UI/UX, jago slicing juga, database juga mantap. paket lengkap ya mas",
    image: "./assets/images/sukarno.jpg",
    rating: 3,
  },
  {
    name: "Bung Tomo",
    quote: "Sepertinya pembuat website ini ingin ku recruit ke perusahaanku",
    image: "./assets/images/bung-tomo.jpg",
    rating: 4,
  },
  {
    name: "R.A. Kartini",
    quote: "Pembuat wbsite andalanku dari dulu hingga nanti",
    image: "./assets/images/kartini.pg.jpg",
    rating: 5,
  },
  {
    name: "Kapiten Pattimura",
    quote:
      "Selalu ada plus minus, tapi sama mas aji plusnya bisa membuat Indonesia merdeka lagi",
    image: "./assets/images/pattimura.png",
    rating: 4,
  },
];

function showTestimonials() {
  let testimonialforHTML = "";

  testimonialData.forEach(function (item) {
    testimonialforHTML += `        
        <div class="testi-card">
         <img src="${item.image}" alt="testi">
         <div class="testi-desc">
             <p class="quotes">"${item.quote}"</p>
             <p class="author">- ${item.name}</p>
             <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
         </div>
     </div>`;
  });

  document.getElementById("testimonials").innerHTML = testimonialforHTML;
}

showTestimonials();

function filterTestimonials(rating) {
  let testimonialforHTML = "";

  const testimonialsFiltered = testimonialData.filter(function (item) {
    return item.rating == rating;
  });

  //   console.log(filterTestimonials);

  if (testimonialsFiltered.length === 0) {
    testimonialforHTML += `<h1>Data not found!</h1>`;
  } else {
    testimonialsFiltered.forEach(function (item) {
      testimonialforHTML += `        
            <div class="testi-card">
             <img src="${item.image}" alt="testi">
             <div class="testi-desc">
                 <p class="quotes">"${item.quote}"</p>
                 <p class="author">- ${item.name}</p>
                 <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
             </div>
         </div>`;
    });
  }

  document.getElementById("testimonials").innerHTML = testimonialforHTML;
}
