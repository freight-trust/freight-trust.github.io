# 

> [Docs](https://docs.rosetta-technology.io/index.html) » [The Roseta
> DSL documentafion](https://docs.rosetta-technology.io/dsl/index.html)
> » Roseta Modelling Components

Rosetta Modelling Components

> **The Roseta syntax can express ﬁve types of model components**:
> 
> Data Annotafion
> 
> Data Validafion (or *condifion*) Funcfion
> 
> Mapping (or *synonym*)
> 
> This documentafion details the purpose and features of each type of
> model component and highlights the relafionships that exist among
> those. As the inifial live applicafion of the Roseta DSL, examples
> from the ISDA CDM will be used to illustrate each of those features.

# Data Component

> **The Roseta DSL provides four data deﬁnifion components** that are
> used to model data, grouped into two pairs:
> 
> Type and Atribute
> 
> Enumerafion and Enumerafion Value

## Type and Attribute

### Purpose

> A *type* describes an *enfity* (also somefimes referred to as an
> *object* or a *class*) in the model and is deﬁned by a plain-text
> descripfion and a set of *atributes*. Atributes specify the granular
> elements composing the enfity.

### Syntax

> The deﬁnifion of a type starts with the keyword , followed by the type
> name. A
> 
> colon punctuafion introduces the rest of the deﬁnifion.
> 
> The ﬁrst component of the deﬁnifion is a plain-text descripfion of the
> type.
> 
> Descripfions use quotafion marks (to mark a string) in between angle
> brackets
> 
> 1.  Descripfions, although not generafing any executable code, are
>     integral meta-data components of the model. As modelling best
>     pracfice, a deﬁnifion ought to exist for every artefact and be
>     clear and comprehensive.
> 
> Then the deﬁnifion of the type lists its component atributes. Each
> atribute is deﬁned by four components, syntacfically ordered as:
> 
> name type
> 
> cardinality: see Cardinality Secfion descripfion
> 
>  **Note**
> 
> The Roseta DSL does not use any delimiter to end deﬁnifions. All model
> 
> deﬁnifions start with a similar opening keyword as , so the start of a
> new
> 
> deﬁnifion marks the end of the previous one. For readability more
> generally, the Roseta DSL looks to eliminate all the delimiters that
> are ohen used in tradifional
> 
> programming languages (such as curly braces or semi-colon ).
> 
> Each atribute can be speciﬁed either as a basic type, a type or an
> enumerafion. The set of basic types available in the Roseta DSL are
> controlled at the language level by
> 
> the deﬁnifion:
> 
> Text - Number - Logic -
> 
> (for integer) and
> 
> (for ﬂoat)
> 
> Date and Time - , and
> 
> The Roseta DSL convenfion is that type names use the *PascalCase*
> (starfing with a capital leter, also referred to as the *upper*
> [CamelCase](https://en.wikipedia.org/wiki/Camel_case)), while atribute
> names use the *camelCase* (starfing with a lower case leter, also
> referred to as the *lower* camelCase). Type names need to be unique
> across the model. All those requirements are controlled by the Roseta
> DSL grammar.

### Time

> For fime zone adjustments, a fime zone qualiﬁer can be speciﬁed
> alongside a fime in one of two ways:
> 
> Through the basic type, which needs to be expressed either as
> 
> [UTC](https://en.wikipedia.org/wiki/Coordinated_Universal_Time) or as
> an oﬀset to UTC, as speciﬁed by the ISO 8601 standard.
> 
> Through the type, where fime is speciﬁed alongside a
> 
> business center. This is used to specify a fime dimension in relafion
> to a future event: e.g. the earliest or latest exercise fime of an
> opfion.
> 
> While there has been discussion as to whether the Roseta DSL should
> support
> 
> dates which are speciﬁed as an oﬀset to UTC with the suﬃx, no posifive
> 
> conclusion has been reached. The main reason is that all dates which
> need a business date context can already specify an associated
> business center.

### Inheritance

> **The Roseta DSL supports an inheritance mechanism**, when a type
> inherits its deﬁnifion and behaviour (and therefore all of its
> atributes) from another type and adds its own set of atributes on top.
> Inheritance is supported by the
> 
> keyword next to the type name.
> 
>  **Note**
> 
> For clarity purposes, the documentafion snippets omit the synonyms and
> deﬁnifions that are associated with the classes and atributes, unless
> the purpose of the snippet is to highlight some of those features.

## Enumeration and Enumeration Value

### Purpose

> **Enumerafion is the mechanism through which an atribute may only take
> some speciﬁc controlled values**. An *enumerafion* is the container
> for the corresponding set of controlled (or enumerafion) values.
> 
> This mimics the *scheme* concept, whose values may be speciﬁed as part
> of an exisfing standard and can be represented through an enumerafion
> in the Roseta DSL. Typically, a scheme with no deﬁned values is
> represented as a basic
> 
> type.

### Syntax

> Enumerafions are very simple modelling containers, which are deﬁned in
> the same way as other model components. The deﬁnifion of an
> enumerafion starts with the
> 
> keyword, followed by the enumerafion name. A colon punctuafion
> 
> introduces the rest of the deﬁnifion, which contains a plain-text
> descripfion of the enumerafion and the list of enumerafion values.
> 
> Enumerafion names must be unique across a model. The Roseta DSL naming
> convenfion is the same as for types and must use the upper CamelCase
> (PascalCase). In addifion the enumerafion name should end with the
> suﬃx Enum.
> 
> Enumerafion values have a restricted syntax to facilitate their
> integrafion with executable code: they cannot start with a numerical
> digit, and the only special
> 
> character that can be associated with them is the underscore .
> 
> In order to handle the integrafion of scheme values which can have
> special characters, the Roseta DSL allows to associate a **display
> name** to any enumerafion
> 
> value. For those enumerafion values, special characters are replaced
> with while
> 
> the entry corresponds to the actual value.
> 
> An example is the day count fracfion scheme for interest rate
> calculafion, which
> 
> includes values such as
> 
> to the
> 
> and and
> 
> 1.  These are associated as enumerafion values, respecfively.

# Annotation Component

## Annotation Deﬁnition

### Purpose

> Annotafions allow to associate meta-informafion to model components,
> which can serve a number of purposes:
> 
> purely syntacfic, to provide addifional guidance when navigafing model
> components
> 
> to add constraints to a model that may be enforced by syntax
> validafion to modify the actual behaviour of a model in generated code
> 
> Examples of annotafions and their usage for diﬀerent purposes are
> illustrated below.

### Syntax

> Annotafions are deﬁned in the same way as other model components. The
> deﬁnifion
> 
> of an annotafion starts with the keyword, followed by the annotafion
> 
> name. A colon punctuafion introduces the rest of the deﬁnifion,
> starfing with a
> 
> plain-text descripfion of the annotafion.
> 
> Annotafion names must be unique across a model. The Roseta DSL naming
> convenfion is to use a (lower) camelCase.
> 
> It is possible to associate atributes to an annotafion (see example),
> even
> 
> though some annotafions may not require any further atribute. For
> instance:

## Meta-Data and Reference

### Purpose

> The
> 
> and atributes.

annotafion allows to associate a set of meta-data qualiﬁers to types

> Each atribute of the annotafion corresponds to a qualiﬁer:
> 
> The meta-data qualiﬁer indicates a type that is referenceable, so that
> a
> 
> unique idenfiﬁer can be associated to objects of that type. This
> allows to replicates the cross-referencing mechanism used in XML to
> provide data integrity
> 
> within the context of an instance document. The
> 
> replicates the
> 
> meta-
> 
> data as used in the FpML standard, which associates a cross-reference
> value to the object’s data source.
> 
> The types.
> 
> meta-data qualiﬁer provides the same funcfionality as
> 
> but for basic
> 
> The meta-data qualiﬁer indicates that the atribute may be speciﬁed
> 
> as a reference, using the of a referenceable object as meta-data. This
> 
> replicates the (for *hyper-text reference*) meta-data as used in the
> FpML
> 
> standard, where the cross-reference value may be speciﬁed as
> meta-informafion in the atribute’s data source.
> 
> The meta-data qualiﬁer speciﬁes a mechanism to control the set of
> 
> values that an atribute can take. The relevant scheme reference may be
> speciﬁed as meta-informafion in the atribute’s data source, so that no
> originafing informafion is disregarded.
> 
> The corresponds to a hash code to be generated by the model
> implementafion.
> 
> The implementafion provided in the Roseta DSL is the de-facto Java
> hash funcfion. It is a *deep hash* that uses the complete set of
> atribute values that compose the type and its atributes, recursively.
> 
>  **Note**
> 
> Some annotafions, such as this metadata qualiﬁcafion, may be provided
> as standard as part of the Roseta DSL itself. Addifional annotafions
> can always be deﬁned for any model.

### Syntax

> Once an annotafion is deﬁned, its name and chosen atribute, if any,
> are used in
> 
> between square brackets to annotate model components. The below
> 
> and types illustrate how meta-data annotafions and their relevant
> 
> atributes can be used in a model:
> 
> A
> 
> the
> 
> qualiﬁer is associated to the type, the
> 
> atribute of type
> 
> type, which means it is referenceable. In qualiﬁer, which is
> associated to the
> 
> , indicates that this atribute can be
> 
> provided as a reference (via its associated key) instead of a copy. An
> example implementafion of this cross-referencing mechanism for these
> types can be found in the Synonym Secfion of the documentafion.

### Partial Key

> Meta-data keys that are generated by a hashing algorithm from an
> object’s atribute values ohen ﬁnd a pracfical use by implementors for
> reconciling and matching data, where equality between hash values is
> considered a proxy for a data match.
> 
> In some cases, it is necessary to remove some of an object’s atribute
> values from the hashing algorithm, when those values are not required
> in the reconciliafion but risk adding noise in the hash that could
> generate false negafives. This is typically the case for meta-data
> qualiﬁers (such as meta-data keys), which may themselves be
> automafically generated by an algorithm. These may result in
> diﬀerences between two objects, even if those objects would have the
> same actual values.
> 
> An implementafion of such parfial key used to be provided as a feature
> of the
> 
> Roseta DSL (with a annotafion). It has now been de-commissioned,
> 
> unfil further evaluafion of its usage emerges that may lead to a
> redesign of this feature.

## Qualiﬁed Type

> The Roseta DSL provides for some special types called *qualiﬁed
> types*, which are speciﬁc to its applicafion in the ﬁnancial domain:
> 
> Calculafion -
> 
> Object qualiﬁcafion -
> 
> Those special types are designed to ﬂag atributes which result from
> running some logic, such that model implementafions can idenfify where
> to stamp the output in the model. The logic is being captured by
> speciﬁc types of funcfions that are detailed in the Funcfion
> Deﬁnifion Secfion.

### Calculation

> The qualiﬁed type, when speciﬁed instead of the type for the atribute,
> 
> represents the outcome of a calculafion. An example usage is the
> conversion from clean price to dirty price for a bond.
> 
> An atribute with the tagged with the funcfion output.
> 
> type is meant to be associated to a funcfion annotafion. The
> atribute’s type is implied by the

### Object Qualiﬁcation

> Similarly,
> 
> and
> 
> represent the outcome of qualiﬁcafion logic
> 
> to infer the type of an object (ﬁnancial product or event) in the
> model. See the atribute, alongside other idenfiﬁer atributes in the
> 
> type:
> 
> Atributes of these types are meant to be associated to an object
> qualiﬁcafion
> 
> funcfion tagged with the annotafion. The annotafion has an atribute
> 
> to specify which type of object (like or ) is being qualiﬁed.
> 
>  **Note**
> 
> The qualiﬁed type feature in the Roseta DSL is under evaluafion and
> may be replaced by a mechanism that is purely based on these funcfion
> annotafions in the future.

# Data Validation Component

> **Data integrity is supported by validafion components that are
> associated to each data type** in the Roseta DSL. There are two types
> of validafion components:
> 
> Cardinality Condifion Statement
> 
> The validafion components associated to a data type generate
> executable code designed to be executed on objects of that type.
> Implementors of the model can use the code generated from these
> validafion components to build diagnosfic tools that can scan objects
> and report on which validafion rules were safisﬁed or broken.
> 
> Typically, the validafion code is included as part of any process that
> creates an object, to verify its validity from the point of creafion.

## Cardinality

> Cardinality is a data integrity mechanism to control how many of each
> atribute an object of a given type can contain. The Roseta DSL borrows
> from XML and speciﬁes
> 
> cardinality as a lower and upper bound in between braces.
> 
> The lower and upper bounds can both be any integer number. A 0 lower
> bound
> 
> means atribute is opfional. A upper bound means an unbounded atribute.
> 
> represents that there must be one and only one atribute of this type.
> When the upper bound is greater than 1, the atribute will be
> considered as a list, to be handled as such in any generated code.
> 
> A separate validafion rule is generated for each atribute’s
> cardinality constraint, so that any cardinality breach can be
> associated back to the speciﬁc atribute and not just to the object
> overall.

## Condition Statement

### Purpose

> *Condifions* are logic statements associated to a data type. They are
> predicates on atributes of objects of that type that evaluate to True
> or False.

### Syntax

> Condifion statements are included in the deﬁnifion of the type that
> they are associated to and are usually appended aher the deﬁnifion of
> the type’s atributes.
> 
> The deﬁnifion of a condifion starts with the keyword, followed by the
> 
> name of the condifion and a colon punctuafion. The condifion’s name
> must be
> 
> unique in the context of the type that it applies to (but does not
> need to be unique across all data types of a given model). The rest of
> the condifion deﬁnifion comprises:
> 
> a plain-text descripfion (opfional)
> 
> a logic expression that applies to the the type’s atributes
> 
> **The Roseta DSL oﬀers a restricted set of language features designed
> to be unambiguous and understandable** by domain experts who are not
> sohware engineers, while minimising unintenfional behaviour. The
> Roseta DSL is not a *Turing-complete* language: it does not support
> looping constructs that can fail (e.g. the loop never ends), nor does
> it nafively support concurrency or I/O operafions. The language
> features that are available in the Roseta DSL to express validafion
> condifions emulate the basic boolean logic available in usual
> programming languages:
> 
> condifional statements: , , boolean operators: ,
> 
> list statements: , , , comparison operators: , , , , ,
> 
>  **Note**
> 
> Condifions are included in the deﬁnifion of the data type that they
> are associated to, so they are “aware” of the context of that data
> type. This is why atributes of that data type can be directly used to
> express the validafion logic, without the need to refer to the type
> itself.

## Special Syntax

> Some speciﬁc language features have been introduced in the Roseta DSL,
> to handle validafion cases where the basic boolean logic components
> would create unecessarily verbose, and therefore less readable,
> expressions. Those use-cases were deemed frequent enough to jusfify
> developing a speciﬁc syntax for them.

### Choice

> Choice rules deﬁne a choice constraint between the set of atributes of
> a type in the Roseta DSL. They allow a simple and robust construct to
> translate the XML *xsd:choicesyntax*, although their usage is not
> limited to those XML use cases.
> 
> The choice constraint can be either:
> 
> **opfional**, represented by the atributes needs to be present, or
> **required**, represented by the atributes needs to be present
> 
> syntax, when at most one of the syntax, when exactly one of the
> 
> While most of the choice rules have two atributes, there is no limit
> to the number of atributes associated with it, within the limit of the
> number of atributes associated with the type.
> 
>  **Note**
> 
> Members of a choice rule need to have their lower cardinality set to
> 0, something which is enforced by a validafion rule.

### One-of (as complement to choice rule)

> In the case where all the atributes of a given type are subject to a
> required choice logic that results in one and only one of them being
> present in any instance of that
> 
> type, the Roseta DSL allows to associate a condifion to the type, as
> short-
> 
> hand to by-pass the implementafion of the corresponding choice rule.
> This feature is illustrated below:

### Only Exists

> The
> 
> component is an adaptafion of the simple
> 
> syntax, that
> 
> veriﬁes that the atribute exists but also that no other atribute of
> the type does.
> 
> This syntax drasfically reduces the condifion expression, which would
> otherwise
> 
> require to combine one
> 
> with mulfiple
> 
> (applied to all other
> 
> atributes). It also makes the logic more robust to future model
> changes, where
> 
> newly introduced atributes would need to be tested for .
> 
>  **Note**
> 
> This condifion is typically applied to atribues of objects whose type
> implements
> 
> a condifion. In this case, the qualiﬁer is redundant with the
> 
> condifion because only one of the atributes can exist. However,
> 
> makes the condifion expression more explicit, and also robust to
> potenfial lihing
> 
> of the condifion.

# Function Component

> **In programming languages, a funcfion is a ﬁxed set of logical
> instrucfions returning an output** which can be parameterised by a set
> of inputs (also known as *arguments*). A funcfion is *invoked* by
> specifying a set of values for the inputs and running the instrucfions
> accordingly. In the Roseta DSL, this type of component has been uniﬁed
> under a single *funcfion* construct.
> 
> Funcfions are a fundamental building block to automate processes,
> because the same set of instrucfions can be executed as many fimes as
> required by varying the inputs to generate a diﬀerent, yet
> determinisfic, result.
> 
> Just like a spreadsheet allows users to deﬁne and make use of
> funcfions to construct complex logic, the Roseta DSL allows to model
> complex processes from reusable funcfion components. Typically,
> complex processes are deﬁned by combining simpler sub-processes, where
> one process’s output can feed as input into another process. Each of
> those processes and sub-processes are represented by a funcfion.
> Funcfions can invoke other funcfions, so they can represent processes
> made up of sub- processes, sub-sub-processes, and so on.
> 
> Reusing small, modular processes has the following beneﬁts:
> 
> **Consistency**. When a sub-process changes, all processes that use
> the sub- process beneﬁt from that single change.
> 
> **Flexibility**. A model can represent any process by reusing exisfing
> sub-processes. There is no need to deﬁne each process explicitly:
> instead, we pick and choose from a set of pre-exisfing building
> blocks.
> 
> Where widely adopted industry processes already exist, they should be
> reused and not redeﬁned. Some examples include:
> 
> Mathemafical funcfions. Funcfions such as sum, absolute, and average
> are widely understood, so do not need to be redeﬁned in the model.
> 
> Reference data. The process of looking-up through reference data is
> usually provided by exisfing industry ufilifies and a model should
> look to re-use it but not re-implement it.
> 
> Quanfitafive ﬁnance. Many quanfitafive ﬁnance solufions, some
> open-source, already deﬁnes granular processes such as:
> 
> compufing a coupon schedule from a set of parameters adjusfing dates
> given a holiday calendar
> 
> This concept of combining and reusing small components is also
> consistent with a modular component approach to modelling.

## Function Speciﬁcation

### Purpose

> **Funcfion speciﬁcafion components are used to deﬁne the processes
> applicable to a domain model** in the Roseta DSL. A funcfion
> speciﬁcafion deﬁnes the funcfion’s inputs and/or output through
> their *types* (or *enumerafions*) in the data model. This amounts to
> specifying the
> [API](https://en.wikipedia.org/wiki/Application_programming_interface)
> that implementors should conform to when building the funcfion that
> supports the corresponding process.
> 
> Standardising those APIs guarantees the integrity, inter-operability
> and consistency of the automated processes supported by the domain
> model.

### Syntax

> Funcfions are deﬁned in the same way as other model components. The
> syntax of a
> 
> funcfion speciﬁcafion starts with the keyword followed by the funcfion
> name.
> 
> A colon punctuafion introduces the rest of the deﬁnifion.
> 
> The Roseta DSL convenfion for a funcfion name is to use a PascalCase
> (upper [CamelCase](https://en.wikipedia.org/wiki/Camel_case)) word.
> The funcfion name needs to be unique across all types of funcfions in
> a model and validafion logic is in place to enforce this.
> 
> The rest of the funcfion speciﬁcafion supports the following
> components:
> 
> plain-text decripfions
> 
> inputs and output atributes (the later is mandatory) condifion
> statements on inputs and output
> 
> output construcfion statements

### Descriptions

> The role of a funcfion must be clear for implementors of the model to
> build applicafions that provide such funcfionality. To beter
> communicate the intent and use of funcfions, Roseta supports mulfiple
> plain-text descripfions in funcfions.
> 
> Descripfions can be provided for the funcfion itself, for any input
> and output and for any statement block.
> 
> Look for occurences of text descripfions in the snippets below.

### Inputs and Output

> Inputs and output are a funcfion’s equivalent of a type’s atributes.
> As in a ,
> 
> each
> 
> atribute is deﬁned by a name, data type (as either a , or
> 
> ) and cardinality.
> 
> At minimum, a funcfion must specify its output atribute, using the
> keyword
> 
> also followed by a colon .
> 
> Most funcfions, however, also require inputs, which are also expressed
> as atributes,
> 
> using the
> 
> keyword.
> 
> is plural whereas
> 
> is singular, because a
> 
> funcfion may only return one type of output but may take several types
> of inputs.

### Conditions

> A funcfion’s inputs and output can be constrained using *condifions*.
> Each condifion is expressed as a logical statement that evaluates to
> True or False, using the same language features as those available to
> express condifion statements in data types and detailed in the
> Condifion Statement Secfion.
> 
> Condifion statements in a funcfion can represent either:
> 
> a **pre-condifion**, using the keyword, applicable to inputs only and
> 
> evaluated prior to execufing the funcfion, or a **post-condifion**,
> using the

keyword, applicable to inputs and

> output and evaluated aher execufing the funcfion (once the output is
> known)
> 
> Condifions are an essenfial feature of the deﬁnifion of a funcfion. By
> constraining the inputs and output, they deﬁne the constraints that
> impementors of this funcfion must safisfy, so that it can be safely
> used for its intended purpose as part of a process.
> 
> func EquityPriceObservation: \<"Function specification for the
> observation of an equity price, based on the attributes of the
> 'EquityValuation' class."\>
> 
> inputs:
> 
> equity Equity (1..1)
> 
> valuationDate AdjustableOrRelativeDate (1..1)
> 
> valuationTime BusinessCenterTime (0..1)
> 
> timeType TimeTypeEnum (0..1)
> 
> determinationMethod DeterminationMethodEnum (1..1) output:
> 
> observation ObservationPrimitive (1..1)
> 
> condition: \<"Optional choice between directly passing a time or a
> timeType, which has to be resolved into a time based on the
> determination method."\>
> 
> **if** valuationTime exists **then** timeType is absent **else if**
> timeType exists **then** valuationTime is absent **else** False
> 
> post-condition: \<"The date and time must be properly resolved as
> attributes on the output."\>
> 
> observation **→** date **=** ResolveAdjustableDate(valuationDate) and
> **if** valuationTime exists **then** observation **→** time **=**
> 
> TimeZoneFromBusinessCenterTime(valuationTime)
> 
> **else** observation **→** time **=**
> ResolveTimeZoneFromTimeType(timeType, determinationMethod)
> 
> post-condition: \<"The number recorded in the observation must match
> the number fetched from the source."\>
> 
> observation **→** observation **=** EquitySpot(equity, observation
> **→** date, observation **→** time)
> 
>  **Note**
> 
> The funcfion syntax intenfionally mimics the type syntax in the Roseta
> DSL regarding the use of descripfions, atributes (inputs and output)
> and condifions, to provide consistency in the expression of model
> deﬁnifions.

## Function Deﬁnition

> **The Roseta DSL allows to further deﬁne the business logic of a
> funcfion**, by building the funcfion output instead of just specifying
> the funcfion’s API. The creafion of valid output objects can be fully
> or parfially deﬁned as part of a funcfion speciﬁcafion, or completely
> leh to the implementor.
> 
> A funcfion is **fully deﬁned** when all validafion constraints on the
> output object have been safisﬁed as part of the funcfion speciﬁcafion.
> In this case, the generated code is directly usable in an
> implementafion.
> 
> A funcfion is **parfially deﬁned** when the output object’s validafion
> constraints are only parfially safisﬁed. In this case, implementors
> will need to extend the generated code and assign the remaining values
> on the output object.
> 
> A funcfion must be applied to a speciﬁc use case in order to determine
> whether it is fully *deﬁned* or *parfially deﬁned*. There are a number
> of fully deﬁned funcfion cases explained in further detail below.
> 
> The Roseta DSL only provides a limited set of language features. To
> build the complete processing logic for a *parfially deﬁned* funcfion,
> model implementors are meant to extend the code generated from the
> Roseta DSL once it is expressed in a fully featured programming
> language. For instance in Java, a funcfion speciﬁcafion generates an
> *interface* that needs to be extended to be executable.
> 
> The output object will be systemafically validated when invoking a
> funcfion, so all funcfions require the output object to be fully valid
> as part of any model implementafion.

### Output Construction

> In the
> 
> example above, the
> 
> statements
> 
> assert whether the observafion’s date and value are correctly
> populated according to the output of other, sub-funcfions, but
> delegates the construcfion of that output to implementors of the
> funcfion.
> 
> In pracfice, implementors of the funcfion can be expected to re-use
> those sub-
> 
> funcfions (
> 
> and
> 
> ) to construct the output. The
> 
> drawback is that those sub-funcfions are likely to be executed twice:
> once to build the output and once to run the validafion.
> 
> For eﬃciency, the funcfion syntax in the Roseta DSL allows to directly
> build the output by assigning its values. Funcfion implementors do not
> have to build those values themselves, because the funcfion already
> provides them by default, so the corresponding post-condifions are
> redundant and can be removed.
> 
> The example above could be rewriten as follows:
> 
> **The Roseta DSL also supports a number of fully deﬁned funcfion
> cases**, where the output is being built up to a valid state:
> 
> Object qualiﬁcafion Calculafion
> 
> Short-hand funcfion
> 
> Those funcfions are typically associated to an annotafion, as
> described in the Qualiﬁed Type Secfion, to instruct code generators to
> create concrete funcfions.

### Object Qualiﬁcation Function

> **The Roseta DSL supports the qualiﬁcafion of ﬁnancial objects from
> their underlying components** according to a given classiﬁcafion
> taxonomy, in order to support a composable model for those objects
> (e.g. ﬁnancial products, legal agreements or their associated
> lifecycle events).
> 
> Object qualiﬁcafion funcfions evaluate a combinafion of asserfions
> that uniquely characterise an input object according to a chosen
> classiﬁcafion. Each funcfion is
> 
> associated to a qualiﬁcafion name (a from that classiﬁcafion) and
> returns a
> 
> boolean. This boolean evaluates to True when the input safisﬁes all
> the criteria to be idenfiﬁed according to that qualiﬁcafion name.
> 
> Object qualiﬁcafion funcfions are associated to a annotafion that
> 
> speciﬁes the type of object being qualiﬁed. The funcfion name start
> with the
> 
> preﬁx, followed by an underscore upper
> [CamelCase](https://en.wikipedia.org/wiki/Camel_case) (PascalCase)
> word, using
> 
> 1.  The naming convenfion is to have an to append granular
>     qualiﬁcafion
> 
> names where the classiﬁcafion may use other types of separators (like
> space or
> 
> colon ).
> 
> Syntax validafion logic based on the this.
> 
> annotafion is in place to enforce

### Calculation Function

> Calculafion funcfions deﬁne a calculafion output that is ohen, though
> not
> 
> exclusively, of type
> 
> 1.  They must end with an
> 
> statement that
> 
> fully deﬁnes the calculafion result.
> 
> Calculafion funcfions are associated to the annotafion.

### Alias

> The funcfion syntax supports the deﬁnifion of *aliases* that are only
> available in the context of the funcfion. Aliases work like temporary
> variable assignments used in programming languages and are
> parficularly useful in fully deﬁned funcfions.
> 
> The above example builds an interest rate calculafion using aliases to
> deﬁne the *calculafion amount*, *rate* and *day count fracfion* as
> temporary variables, and ﬁnally assigns the *ﬁxed amount* output as
> the product of those three variables.

### Short-Hand Function

> Short-hand funcfions are funcfions which are designed to provide a
> compact syntax for operafions that need to be frequently invoked in
> the model - for instance, model indirecfions when the corresponding
> model expression may be deemed too long or cumbersome:
> 
> which could be invoked as part of mulfiple other funcfions that use
> the object by simply stafing:

# Mapping Component

## Synonym

### Purpose

> *Synonym* is the baseline building block to map a model expressed in
> the Roseta DSL to alternafive data representafions, whether those are
> open standards or proprietary. Synonyms can be complemented by mapping
> logic when the relafionship is not a one-to-one or is condifional.
> 
> Synonyms are speciﬁed at the atribute level for a data type. Synonyms
> can also be associated to enumerafions and are speciﬁed at the
> enumerafion value level.
> 
> Mappings are typically implemented by traversing the model tree down,
> so knowledge of the context of an atribute (i.e. the type in which it
> is used) determines what it should map to. Knowledge about the
> upper-level type would be lost if synonyms were implemented at the
> class level.
> 
> There is no limit to the number of synonyms that can be associated to
> any atribute, and there can even be several synonyms for a given data
> source (e.g. in the case of a condifional mapping).

### Syntax

> Synonyms are introduced by the atribute in between square brackets
> synonym syntax has two components:
> 
> keyword and are speciﬁed for each
> 
> , same as an annotafion. The baseline
> 
> **source**, which possible values are controlled by a special
> enumerafion
> 
> type of
> 
> **value**, which is a
> 
> the source
> 
> For example for a data type:
> 
> that idenfiﬁes the name of the atribute as it is found in
> 
> Or an enumerafion:
> 
> enum NaturalPersonRoleEnum: \<"The enumerated values for the natural
> person’s role."\>
> 
> Broker \<"The person who arranged with a client to execute the
> trade."\> \[synonym FpML\_5\_10 , CME\_SubmissionIRS\_1\_0 ,
> CME\_ClearedConfirm\_1\_17 value
> 
> "Broker"\]
> 
> Buyer \<"Acquirer of the legal title to the financial instrument."\>
> \[synonym FpML\_5\_10, CME\_SubmissionIRS\_1\_0,
> CME\_ClearedConfirm\_1\_17 value
> 
> "Buyer"\]
> 
> DecisionMaker \<"The party or person with legal responsibility for
> authorization of the execution of the transaction."\>
> 
> \[synonym FpML\_5\_10, CME\_SubmissionIRS\_1\_0,
> CME\_ClearedConfirm\_1\_17 value "DecisionMaker"\]
> 
> ExecutionWithinFirm \<"Person within the firm who is responsible for
> execution of the transaction."\>
> 
> \[synonym FpML\_5\_10, CME\_SubmissionIRS\_1\_0,
> CME\_ClearedConfirm\_1\_17 value "ExecutionWithinFirm"\]
> 
> InvestmentDecisionMaker \<"Person who is responsible for making the
> investment decision."\>
> 
> \[synonym FpML\_5\_10, CME\_SubmissionIRS\_1\_0,
> CME\_ClearedConfirm\_1\_17 value "InvestmentDecisionMaker"\]
> 
> Seller \<"Seller of the legal title to the financial instrument."\>
> \[synonym FpML\_5\_10, CME\_SubmissionIRS\_1\_0,
> CME\_ClearedConfirm\_1\_17 value
> 
> "Seller"\]
> 
> Trader \<"The person who executed the trade."\>
> 
> \[synonym FpML\_5\_10, CME\_SubmissionIRS\_1\_0,
> CME\_ClearedConfirm\_1\_17 value "Trader"\]
> 
>  **Note**
> 
> The synonym value is of type
> 
> to facilitate integrafion with executable
> 
> code. The alternafive approach consisfing of specifying the value as a
> compafible idenfiﬁer alongside a display name has been disregarded
> because it has been deemed not appropriate to create a “code-friendly”
> value for the respecfive synonyms.
> 
> A further set of atributes can be associated with a synonym, to
> address speciﬁc use cases:
> 
> **path** to allow mapping when data is nested in mulfiple levels
> within the respecfive model.
> 
> **hint** to allow mapping when data is nested in diﬀerent ways between
> the respecfive models.
> 
> The type provides a good illustrafion of such cases:
> 
> **type** Price: \<"Generic description of the price concept applicable
> across product types, which can be expressed in a number of ways other
> than simply cash price"\>
> 
> cashPrice CashPrice (0..1) \<"Price specified in cash terms, e.g. for
> securities proceeds or fee payment in a contractual product."\>
> 
> \[synonym FpML\_5\_10 value "initialPrice" path "rateOfReturn",
> "underlyer"\] \[synonym FpML\_5\_10 hint "paymentAmount"\]
> 
> \[synonym FpML\_5\_10 hint "fixedAmount"\]
> 
> exchangeRate ExchangeRate (0..1) \<"Price specified as an exchange
> rate between two currencies."\>
> 
> \[synonym FpML\_5\_10 value "exchangeRate"\]
> 
> fixedInterestRate FixedInterestRate (0..1) \<"Price specified as a
> fixed interest rate."\>
> 
> \[synonym FpML\_5\_10, CME\_SubmissionIRS\_1\_0,
> CME\_ClearedConfirm\_1\_17 value "fixedRateSchedule" path
> "calculationPeriodAmount→calculation"\]
> 
> \[synonym FpML\_5\_10, CME\_SubmissionIRS\_1\_0,
> CME\_ClearedConfirm\_1\_17 value "fixedAmountCalculation"\]
> 
> \[synonym FpML\_5\_10, CME\_SubmissionIRS\_1\_0,
> CME\_ClearedConfirm\_1\_17 value "fixedRateSchedule"\]
> 
> \[synonym FpML\_5\_10, CME\_SubmissionIRS\_1\_0,
> CME\_ClearedConfirm\_1\_17 hint "fixedRate"\]
> 
> floatingInterestRate FloatingInterestRate (0..1) \<"Price specified as
> a spread on top of a floating interest rate."
> 
> \[synonym FpML\_5\_10, CME\_SubmissionIRS\_1\_0,
> CME\_ClearedConfirm\_1\_17 value "floatingRateCalculation" path
> "calculationPeriodAmount→calculation"\]
> 
> \[synonym FpML\_5\_10, CME\_SubmissionIRS\_1\_0,
> CME\_ClearedConfirm\_1\_17 value "floatingRateCalculation" path
> "interestCalculation"\]
> 
> \[synonym FpML\_5\_10, CME\_SubmissionIRS\_1\_0,
> CME\_ClearedConfirm\_1\_17 value "floatingRateCalculation"\]
> 
> \[synonym FpML\_5\_10, CME\_SubmissionIRS\_1\_0,
> CME\_ClearedConfirm\_1\_17 value "floatingAmountCalculation"\]
> 
> **tag** or a **componentID** to properly reﬂect the FIX standard,
> which uses those two components. There are only limited examples of
> such at present, as a result of the scope focus on post-execufion use
> cases hence the limited reference to the FIX standard.
> 
> **deﬁnifion** to provide a more explicit reference to the FIX
> enumerafion values which are speciﬁed through a single digit or leter
> posifioned as a preﬁx to the associated deﬁnifion.

### Meta-Data Mapping

> When meta-data are associated to an atribute, as decribed in the
> Meta-Data and Reference Secfion, addifional synonym syntax allows to
> specify how to retrieve the corresponding meta-data from the source.
> This is illustrated by the usage of the
> 
> synonym syntax in the example below:
> 
> The
> 
> atribute has an associated
> 
> 1.  The scheme can be retrieved using
> 
> the source.
> 
> meta-data that is atached to the
> 
> value in the synonym
> 
> To be able to specify an atribute as a reference from an exisfing
> source, the source itself must implement some cross-referencing
> mechanism so that the reference can
> 
> be idenfiﬁed, as in the works as follows:
> 
> / mechanism used in XML. The cross-referencing
> 
> the atribute must specify the idenfiﬁer value for the reference in the
> synonym
> 
> source. For the meta-data of the
> 
> atribute above, this is speciﬁed as the value in the source.
> 
> an idenfiﬁer value must be associated to the object being referenced.
> For the
> 
> type, this is speciﬁed as the shown below:
> 
> meta-data in the synonym source, as
> 
> The below JSON extract illustrates an implementafion of these
> meta-data in the context of a *transacfion event*, which idenfiﬁes the
> parfies to the transacfions as well as the *issuer* of the event (i.e.
> who submits the transacfion message).
> 
> There are two parfies to the event, associated with idenfiﬁers as
> 
> “party1” and “party2”. Their actual are speciﬁed through an FpML
> 
> values are “Party 1” and “Party 2”, which referred to in meta-data.
> Roseta also
> 
> associates an internal meta-data.
> 
> hash to each party, as implementafion of the
> 
> Thanks to the
> 
> qualiﬁer, the
> 
> atribute can simply
> 
> reference the event issuer party as “Party 2” rather than duplicafing
> its components. The cross-reference is sourced from the original FpML
> document using the
> 
> implemented
> 
> synonym. The internal
> 
> points to the
> 
> hash while the points to the “party2” ,
> 
> as sourced from the original FpML document. Also note that the
> 
> itself has an associated
> 
> meta-data by default since its
> 
> class
> 
> has a qualiﬁer.
> 
>  **Note**
> 
> This example is not part of the Roseta DSL but corresponds to the
> default JSON implementafion of the model. The choice of either
> maintaining or shredding external references (such as “party2”), once
> cross-reference has been established using the source document, is up
> to implementors of the model.

## Mapping Rule

### Purpose

> There are cases where the mapping between exisfing standards and
> protocols and their relafion to the model is not one-to-one or is
> condifional. Synonyms have been complemented with a syntax to express
> mapping logic that provides a balance between ﬂexibility and
> legibility.

### Syntax

> The mapping rule syntax diﬀers from the normal Roseta DSL syntax in
> that it is not
> 
> expressed as a stand-alone block with a qualiﬁer preﬁx such as .
> Instead,
> 
> the mapping rule is posifioned as an extension of the synonym syntax.
> Several mapping rule expressions can be associated with a given
> synonym.
> 
> A mapping rule is composed of two (opfional) expressions:
> 
> **mapping value** preﬁxed with , which speciﬁes the value that the
> atribute
> 
> should be set to when the condifional expression is true
> 
> **condifional expression** preﬁxed with mapping value
> 
> , to associate condifional logic to the
> 
> The mapping logic associated with the party role example below
> provides a good illustrafion of such logic:
