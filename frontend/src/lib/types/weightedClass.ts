import type { IGrade } from "./grade"

export interface IWeightedClass {
    classCode: string
    totalWeight: number
    finalWeightedGrade: number
    gradeList: Array<IGrade>
}