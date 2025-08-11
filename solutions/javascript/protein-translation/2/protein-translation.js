//
// This is only a SKELETON file for the 'Protein Translation' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const translate = function () {
  const codonStr = [...arguments][0];
  if (!codonStr) {
    return [];
  }
  const stopCodons = ["UAA", "UAG", "UGA"];
  let codonsArr = [];
  for (let i = 0; i <= codonStr.length - 1; i = i + 3) {
    const codon = codonStr.substr(i, 3);
    if (stopCodons.includes(codon)) {
      break;
    } else {
      codonsArr.push(codon);
    }
  }
  if (!codonsArr.length) {
    return [];
  }
  const aminoAcids = [];
  const codonAminoArr = [
    ["Methionine", "AUG" ],
    ["Phenylalanine", "UUU", "UUC" ],
    ["Leucine", "UUA", "UUG", ],
    ["Serine", "UCU", "UCC", "UCA", "UCG"],
    ["Tyrosine", "UAU", "UAC"],
    ["Cysteine", "UGU", "UGC"],
    ["Tryptophan", "UGG"],
  ];
  for(let codon of codonsArr){
  	let found = false;
  	for(let groupIndex = 0; groupIndex<=codonAminoArr.length - 1; groupIndex++){
  		const [aminoAcid, ...codons] = codonAminoArr[groupIndex];
  		if(codons.includes(codon)){
  			aminoAcids.push(aminoAcid)
  			break;
  		}
  		if(groupIndex===codonAminoArr.length -1 && !found){
  			throw new Error("Invalid codon")
  		}
  	}
  }
  return aminoAcids;
};
