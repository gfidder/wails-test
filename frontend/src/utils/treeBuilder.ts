/* eslint-disable security/detect-object-injection */
import { oidstorage } from "../../wailsjs/go/models";

export interface OidTree {
  name: string;
  oid: string;
  children?: Array<OidTree>;
}

export class TreeSorter {
  #newOids: Array<oidstorage.Oid>;

  constructor(newOids: Array<oidstorage.Oid>) {
    this.#newOids = newOids;
  }

  createOidTree(): OidTree {
    const tree: OidTree = { name: "", oid: "" };
    let i: number;

    for (i = 0; i < this.#newOids.length; i++) {
      const oid = this.#newOids[i];
      if (oid.name === "iso") {
        tree.name = oid.name;
        tree.oid = oid.oid;
        break;
      }
    }

    this.#newOids.splice(i, 1);

    // call some function here?

    const parent = tree;
    const nextChild = tree.children;
    let elementsFound = true;

    while (elementsFound) {
      const parentOidNum = TreeSorter.getNumber(parent.oid);

      const childIndexArray = Array<number>();

      for (let i = 0; i < this.#newOids.length; i++) {
        const oid = this.#newOids[i];
        const oidNumber = TreeSorter.getNumber(oid.oid);
        if (parentOidNum + 1 === oidNumber) {
          childIndexArray.push(i);
        }
      }

      childIndexArray.forEach((index) => {
        const newOidObject: OidTree = {
          name: this.#newOids[index].name,
          oid: this.#newOids[index].oid,
        };
        nextChild?.push(newOidObject);
      });

      childIndexArray.forEach((index) => {
        this.#newOids.splice(index, 1);
      });

      if (this.#newOids.length === 0) {
        elementsFound = false;
        // need something with parent here
      }

      elementsFound = false;
    }

    return tree;
  }

  static isDirectChild(parentOid: string, childOid: string): boolean {
    let isChild = false;

    if (
      TreeSorter.getNumber(parentOid) + 1 ===
      TreeSorter.getNumber(childOid)
    ) {
      const lastPeriodIndex = childOid.lastIndexOf(".");
      const childComparer = childOid.substring(0, lastPeriodIndex);
      console.log(childComparer);

      if (childComparer === parentOid) {
        isChild = true;
      }
    }

    return isChild;
  }

  static getNumber(oid: string): number {
    type Subtraction = 0 | 1;

    let leadingSubtraction: Subtraction = 1;
    const leadingPeriod = oid.at(0) === ".";
    if (!leadingPeriod) {
      leadingSubtraction = 0;
    }

    const splitOids = oid.split(".");
    return splitOids.length - leadingSubtraction;
  }
}
