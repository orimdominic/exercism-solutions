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
  const aminoSeqMapArr = [
    {aminoAcid: "Methionine", sequences: ["AUG"]},
    {aminoAcid: "Phenylalanine", sequences: ["UUU", "UUC"]},
    {aminoAcid: "Leucine", sequences: [ "UUA", "UUG"]},
    {aminoAcid: "Serine", sequences: ["UCU", "UCC", "UCA", "UCG"]},
    {aminoAcid: "Tyrosine", sequences: ["UAU", "UAC"]},
    {aminoAcid: "Cysteine", sequences: ["UGU", "UGC"]},
    {aminoAcid: "Tryptophan", sequences: ["UGG"]},
  ];
  for(let codon of codonsArr){
  	let found = false;
  	for(let groupIndex = 0; groupIndex<=aminoSeqMapArr.length - 1; groupIndex++){
  		const {aminoAcid, sequences} = aminoSeqMapArr[groupIndex];
  		if(sequences.includes(codon)){
  			aminoAcids.push(aminoAcid)
  			break;
  		}
  		if(groupIndex===aminoSeqMapArr.length -1 && !found){
  			throw new Error("Invalid codon")
  		}
  	}
  }
  return aminoAcids;
};
